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
	"github.com/rabbitstack/fibratus/cmd/fibratus/common"
	"github.com/rabbitstack/fibratus/pkg/aggregator"
	"github.com/rabbitstack/fibratus/pkg/alertsender"
	"github.com/rabbitstack/fibratus/pkg/api"
	"github.com/rabbitstack/fibratus/pkg/config"
	"github.com/rabbitstack/fibratus/pkg/filament"
	"github.com/rabbitstack/fibratus/pkg/filter"
	"github.com/rabbitstack/fibratus/pkg/handle"
	"github.com/rabbitstack/fibratus/pkg/kstream"
	"github.com/rabbitstack/fibratus/pkg/ps"
	"github.com/rabbitstack/fibratus/pkg/util/multierror"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var runCmd = &cobra.Command{
	Use:     "run [filter]",
	Short:   "Bootstrap fibratus or a filament",
	Aliases: []string{"start"},
	RunE:    run,
	Example: `
	# Run without the filter
	fibratus run

	# Run with the filter that drops all but events produced by the svchost.exe process
	fibratus run ps.name = 'svchost.exe'

	# Run with the filter that traps all events that were generated by process that contains the 'svc' string and it was started by 'SYSTEM' or 'admin' users
	fibratus run ps.name contains 'svc' and ps.sid in ('NT AUTHORITY\\SYSTEM', 'ARCHRABBIT\\admin')

	# Run the top_keys filament
	fibratus run -f top_keys
	`,
}

var (
	// the run command config
	cfg = config.NewWithOpts(config.WithRun())
)

func init() {

	// initialize flags
	cfg.MustViperize(runCmd)

}

func run(cmd *cobra.Command, args []string) error {
	// initialize config and logger
	if err := common.Init(cfg, true); err != nil {
		return err
	}

	// set up the signals
	stopCh := common.Signals()

	// initialize kernel trace controller and try to start the trace
	ktracec := kstream.NewKtraceController(cfg.Kstream)
	err := ktracec.StartKtrace()
	if err != nil {
		return err
	}

	// bootstrap essential components, including handle, process snapshotters
	// and the kernel stream consumer that will actually collect all the events
	hsnap := handle.NewSnapshotter(cfg, nil)
	psnap := ps.NewSnapshotter(hsnap, cfg)
	kstreamc := kstream.NewConsumer(ktracec, psnap, hsnap, cfg)
	// build the filter from the CLI argument. If we got a valid expression the filter
	// is linked to the kernel stream consumer so it can drop any events that don't match
	// the filter criteria
	kfilter, err := filter.NewFromCLI(args, cfg)
	if err != nil {
		return err
	}
	if kfilter != nil {
		kstreamc.SetFilter(kfilter)
	}
	log.Infof("bootstrapping with pid %d", os.Getpid())

	// user can either instruct to bootstrap a filament or start a regular run. We'll setup
	// the corresponding components accordingly to what we got from the CLI options. If a filament
	// was given, we'll assign it the previous filter if it wasn't provided in the filament init function.
	// Finally, we open the kernel stream flow and run the filament i.e. Python main thread in a new goroutine.
	// In case of a regular run, we additionally setup the aggregator. The aggregator will grab the events
	// from the queue, assemble them into batches and hand over to output sinks.
	var f filament.Filament
	filamentName := cfg.Filament.Name
	if filamentName != "" {
		f, err = filament.New(filamentName, psnap, hsnap, cfg)
		if err != nil {
			return err
		}
		if f.Filter() != nil {
			kstreamc.SetFilter(f.Filter())
		}
		err = kstreamc.OpenKstream()
		if err != nil {
			return multierror.Wrap(err, ktracec.CloseKtrace())
		}
		// load alert senders so emitting alerts is possible from filaments
		err = alertsender.LoadAll(cfg.Alertsenders)
		if err != nil {
			log.Warnf("couldn't load alertsenders: %v", err)
		}
		go func() {
			err = f.Run(kstreamc.Events(), kstreamc.Errors())
			if err != nil {
				log.Error(err)
				stopCh <- struct{}{}
			}
		}()
	} else {
		err = kstreamc.OpenKstream()
		if err != nil {
			return multierror.Wrap(err, ktracec.CloseKtrace())
		}
		// setup the aggregator that forwards events to outputs
		agg, err := aggregator.NewBuffered(
			kstreamc.Events(),
			kstreamc.Errors(),
			cfg.Aggregator,
			cfg.Output,
			cfg.Transformers,
			cfg.Alertsenders,
		)
		if err != nil {
			return err
		}
		defer func() {
			if err := agg.Stop(); err != nil {
				log.Error(err)
			}
		}()
	}

	defer func() {
		_ = ktracec.CloseKtrace()
		_ = kstreamc.CloseKstream()
	}()

	// start the HTTP server
	if err := api.StartServer(cfg); err != nil {
		return err
	}

	<-stopCh

	// shutdown everything gracefully
	if f != nil {
		if err := f.Close(); err != nil {
			return err
		}
	}
	if err := handle.CloseTimeout(); err != nil {
		return err
	}
	if err := api.CloseServer(); err != nil {
		return err
	}

	return nil
}
