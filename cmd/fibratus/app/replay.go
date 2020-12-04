/*
 * Copyright 2019-2020 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package app

import (
	"context"
	"github.com/rabbitstack/fibratus/pkg/aggregator"
	"github.com/rabbitstack/fibratus/pkg/api"
	"github.com/rabbitstack/fibratus/pkg/config"
	"github.com/rabbitstack/fibratus/pkg/filament"
	"github.com/rabbitstack/fibratus/pkg/filter"
	"github.com/rabbitstack/fibratus/pkg/kcap"
	"github.com/rabbitstack/fibratus/pkg/outputs"
	logger "github.com/rabbitstack/fibratus/pkg/util/log"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

var replayCmd = &cobra.Command{
	Use:   "replay",
	Short: "Replay kernel event flow from the kcap file",
	RunE:  replay,
}

var replayConfig = config.NewWithOpts(config.WithReplay())

func init() {
	replayConfig.MustViperize(replayCmd)
}

func replay(cmd *cobra.Command, args []string) error {
	if err := replayConfig.TryLoadFile(replayConfig.File()); err != nil {
		return err
	}
	if err := replayConfig.Init(); err != nil {
		return err
	}
	if err := replayConfig.Validate(); err != nil {
		return err
	}
	if err := logger.InitFromConfig(replayConfig.Log); err != nil {
		return err
	}
	kfilter, err := filter.NewFromCLI(args)
	if err != nil {
		return err
	}
	// initialize kcap reader and try to recover the snapshotters
	// from the captured state
	reader, err := kcap.NewReader(replayConfig.KcapFile, replayConfig)
	if err != nil {
		return err
	}
	hsnap, psnap, err := reader.RecoverSnapshotters()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	filamentConfig := replayConfig.Filament
	filamentName := filamentConfig.Name
	// we don't need the aggregator is user decided to replay the
	// kcap on the filament. Otwherise, we setup the full-fledged
	// buffered aggregator
	var agg *aggregator.BufferedAggregator

	if filamentName != "" {
		f, err := filament.New(filamentName, psnap, hsnap, filamentConfig)
		if err != nil {
			return err
		}
		if f.Filter() != nil {
			kfilter = f.Filter()
		}
		reader.SetFilter(kfilter)

		// returns the channel where events are read from the kcap
		kevents, errs := reader.Read(ctx)

		go func() {
			defer f.Close()
			err = f.Run(kevents, errs)
			if err != nil {
				sig <- os.Interrupt
			}
		}()
	} else {
		if kfilter != nil {
			reader.SetFilter(kfilter)
		}

		// use the channels where events are read from the kcap as aggregator source
		kevents, errs := reader.Read(ctx)

		var err error
		agg, err = aggregator.NewBuffered(
			kevents,
			errs,
			replayConfig.Aggregator,
			outputs.Config{Type: replayConfig.Output.Type, Output: replayConfig.Output.Output},
			replayConfig.Transformers,
			replayConfig.Alertsenders,
		)
		if err != nil {
			return err
		}
	}
	// start the HTTP server
	if err := api.StartServer(replayConfig); err != nil {
		return err
	}
	signal.Notify(sig, os.Kill, os.Interrupt)
	<-sig
	// stop reader consumer goroutines
	cancel()

	if agg != nil {
		if err := agg.Stop(); err != nil {
			return err
		}
	}

	return nil
}
