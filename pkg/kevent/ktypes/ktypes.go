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

package ktypes

import (
	"hash/fnv"
	"syscall"
)

// Ktype identifies a kernel event type. It comprises the kernel event GUID + one extra opcode byte to uniquely identify a kernel event
type Ktype [17]byte

var (
	// CreateProcess identifies process creation kernel events
	CreateProcess = Pack(syscall.GUID{Data1: 0x3d6fa8d0, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x0, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}, 1)
	// TerminateProcess identifies process termination kernel events
	TerminateProcess = Pack(syscall.GUID{Data1: 0x3d6fa8d0, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x0, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}, 2)
	// EnumProcess represents the start data collection process event that enumerates processes that are currently running at the time the kernel session starts
	EnumProcess = Pack(syscall.GUID{Data1: 0x3d6fa8d0, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x0, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}, 3)

	// CreateThread identifies thread creation kernel events
	CreateThread = Pack(syscall.GUID{Data1: 0x3d6fa8d1, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x0, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}, 1)
	// TerminateThread identifies thread termination kernel events
	TerminateThread = Pack(syscall.GUID{Data1: 0x3d6fa8d1, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x0, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}, 2)
	// EnumThread represents the start data collection thread event that enumerates threads that are currently running at the time the kernel session starts
	EnumThread = Pack(syscall.GUID{Data1: 0x3d6fa8d1, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x0, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}, 3)

	// FileRundown events are generated by kernel rundown logger to enumerate all open files on the start of the kernel session
	FileRundown = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 36)
	// CreateFile represents events that create/open a file or I/O device
	CreateFile = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 64)
	// ReleaseFile represents events that occur when the last file handle is disposed
	ReleaseFile = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 65)
	// CloseFile represents events that dispose existing kernel file objects
	CloseFile = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 66)
	// ReadFile represents events that read data from the file or I/O device
	ReadFile = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 67)
	// WriteFile represents events that write data to the file or I/O device
	WriteFile = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 68)
	// SetFileInformation represents events that set file information
	SetFileInformation = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 69)
	// DeleteFile identifies file deletion events
	DeleteFile = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 70)
	// RenameFile identifies events that are responsible for renaming files
	RenameFile = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 71)
	// EnumDirectory identifies enumerate directory and directory notification events
	EnumDirectory = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 72)
	// FileOpEnd signals the finalization of the file operation
	FileOpEnd = Pack(syscall.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}, 76)

	// RegCreateKey represents registry key creation kernel events
	RegCreateKey = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 10)
	// RegOpenKey represents registry open key kernel events
	RegOpenKey = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 11)
	// RegDeleteKey represents registry key deletion kernel events
	RegDeleteKey = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 12)
	// RegQueryValue represents registry query key kernel events
	RegQueryKey = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 13)
	// RegSetValue represents registry set value kernel events
	RegSetValue = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 14)
	// RegDeleteValue are kernel events for registry value removals
	RegDeleteValue = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 15)
	// RegQueryValue are kernel events for registry value queries
	RegQueryValue = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 16)
	// RegCreateKCB represents kernel events for KCB (Key Control Block) creation requests
	RegCreateKCB = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 22)
	// RegDeleteKCB represents kernel events for KCB(Key Control Block) closures
	RegDeleteKCB = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 23)
	// RegKCBRundown enumerates the registry keys open at the start of the kernel session.
	RegKCBRundown = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 25)
	// RegOpenKeyV1 represents registry open key kernel event. It looks like this event type defines identical kernel event type as RegOpenKey
	RegOpenKeyV1 = Pack(syscall.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}, 27)

	// UnloadImage represents unload image kernel events
	UnloadImage = Pack(syscall.GUID{Data1: 0x2cb15d1d, Data2: 0x5fc1, Data3: 0x11d2, Data4: [8]byte{0xab, 0xe1, 0x0, 0xa0, 0xc9, 0x11, 0xf5, 0x18}}, 2)
	// EnumImage represents kernel events that is triggered to enumerate all loaded images
	EnumImage = Pack(syscall.GUID{Data1: 0x2cb15d1d, Data2: 0x5fc1, Data3: 0x11d2, Data4: [8]byte{0xab, 0xe1, 0x0, 0xa0, 0xc9, 0x11, 0xf5, 0x18}}, 3)
	// LoadImage represents load image kernel events that are triggered when a DLL or executable file  is loaded
	LoadImage = Pack(syscall.GUID{Data1: 0x2cb15d1d, Data2: 0x5fc1, Data3: 0x11d2, Data4: [8]byte{0xab, 0xe1, 0x0, 0xa0, 0xc9, 0x11, 0xf5, 0x18}}, 10)

	// AcceptTCPv4 represents the TCPv4 kernel events for accepting connection requests from the socket queue.
	AcceptTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 15)
	// AcceptTCPv6 represents the TCPv6 kernel events for accepting connection requests from the socket queue.
	AcceptTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 31)
	// SendTCPv4 represents the TCPv4 kernel events for sending data to the connected socket.
	SendTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 10)
	// SendTCPv6 represents the TCPv6 kernel events for sending data to the connected socket.
	SendTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 26)
	// SendUDPv4 represents the UDPv4 kernel events for sending datagrams to connectionless sockets.
	SendUDPv4 = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xc0, 0xb7, 0xc8, 0x0f, 0x80}}, 10)
	// SendUDPv6 represents the UDPv6 kernel events for sending datagrams to connectionless sockets.
	SendUDPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 26)

	RecvTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 11)
	RecvTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 27)
	RecvUDPv4 = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xc0, 0xb7, 0xc8, 0x0f, 0x80}}, 10)
	RecvUDPv6 = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xc0, 0xb7, 0xc8, 0x0f, 0x80}}, 27)

	ConnectTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 12)
	ConnectTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 28)

	DisconnectTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 13)
	DisconnectTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 29)
	Disconnect      = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 42)

	ReconnectTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 16)
	ReconnectTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 32)

	RetransmitTCPv4 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 14)
	RetransmitTCPv6 = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 30)

	// Accept represents the global kernel event type for both TCP v4/v6 connections. Note this is an artificial kernel event that is never published by the provider.
	Accept = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 46)
	// Send represents the global kernel event for all variants of sending data to sockets. Note this is an artificial kernel event that is never published by the provider.
	Send       = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xa9c9, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 72)
	Recv       = Pack(syscall.GUID{Data1: 0xbf3a50c5, Data2: 0xc8e0, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xc0, 0xb7, 0xc8, 0x0f, 0x80}}, 75)
	Connect    = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 40)
	Reconnect  = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 47)
	Retransmit = Pack(syscall.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}, 44)

	// CreateHandle represents handle creation kernel event
	CreateHandle = Pack(syscall.GUID{Data1: 0x89497f50, Data2: 0xeffe, Data3: 0x4440, Data4: [8]byte{0x8c, 0xf2, 0xce, 0x6b, 0x1c, 0xdc, 0xac, 0xa7}}, 32)
	// CloseHandle represents handle closure kernel event
	CloseHandle = Pack(syscall.GUID{Data1: 0x89497f50, Data2: 0xeffe, Data3: 0x4440, Data4: [8]byte{0x8c, 0xf2, 0xce, 0x6b, 0x1c, 0xdc, 0xac, 0xa7}}, 33)

	// UnknownKtype designates unknown kernel event type
	UnknownKtype = Pack(syscall.GUID{}, 0)
)

// String returns the string representation of the kernel event type. If event is unknown a GUID representation that includes the GUID of the event's provider + the opcode type is presented.
func (k Ktype) String() string {
	switch k {
	case CreateProcess:
		return "CreateProcess"
	case TerminateProcess:
		return "TerminateProcess"
	case CreateThread:
		return "CreateThread"
	case TerminateThread:
		return "TerminateThread"
	case CreateFile:
		return "CreateFile"
	case CloseFile:
		return "CloseFile"
	case ReleaseFile:
		return "ReleaseFile"
	case ReadFile:
		return "ReadFile"
	case WriteFile:
		return "WriteFile"
	case SetFileInformation:
		return "SetFileInformation"
	case DeleteFile:
		return "DeleteFile"
	case RenameFile:
		return "RenameFile"
	case EnumDirectory:
		return "EnumDirectory"
	case FileOpEnd:
		return "FileOpEnd"
	case FileRundown:
		return "FileRundown"
	case CreateHandle:
		return "CreateHandle"
	case CloseHandle:
		return "CloseHandle"
	case RegKCBRundown:
		return "RegKCBRundown"
	case RegOpenKey, RegOpenKeyV1:
		return "RegOpenKey"
	case RegCreateKey:
		return "RegCreateKey"
	case RegDeleteKey:
		return "RegDeleteKey"
	case RegDeleteValue:
		return "RegDeleteValue"
	case RegQueryKey:
		return "RegQueryKey"
	case RegQueryValue:
		return "RegQueryValue"
	case RegCreateKCB:
		return "RegCreateKCB"
	case LoadImage:
		return "LoadImage"
	case UnloadImage:
		return "UnloadImage"
	case Accept:
		return "Accept"
	case Send:
		return "Send"
	case Recv:
		return "Recv"
	case Connect:
		return "Connect"
	case Reconnect:
		return "Reconnect"
	case Disconnect:
		return "Disconnect"
	case Retransmit:
		return "Retransmit"
	default:
		return string(k[:])
	}
}

// Hash calculates the hash number of the ktype.
func (k Ktype) Hash() uint32 {
	h := fnv.New32()
	_, err := h.Write([]byte(k.String()))
	if err != nil {
		return 0
	}
	return h.Sum32()
}

// Exists determines whether particular ktype exists.
func (k Ktype) Exists() bool {
	switch k {
	case EnumProcess, EnumThread, FileRundown, FileOpEnd, ReleaseFile, EnumImage, RegCreateKCB, RegKCBRundown:
		return true
	default:
		// for composite kernel events we match against a single global ktype. This way
		// we use a unique kernel type to group several kernel events. For example, `Send`
		// designates all network Send regardless of transport protocol or IP version
		if k == AcceptTCPv4 || k == AcceptTCPv6 {
			return true
		}
		if k == ConnectTCPv4 || k == ConnectTCPv6 {
			return true
		}
		if k == ReconnectTCPv4 || k == ReconnectTCPv6 {
			return true
		}
		if k == RetransmitTCPv4 || k == RetransmitTCPv6 {
			return true
		}
		if k == DisconnectTCPv4 || k == DisconnectTCPv6 {
			return true
		}
		if k == SendTCPv4 || k == SendTCPv6 || k == SendUDPv4 || k == SendUDPv6 {
			return true
		}
		if k == RecvTCPv4 || k == RecvTCPv6 || k == RecvUDPv4 || k == RecvUDPv6 {
			return true
		}
		_, ok := kevents[k]
		return ok
	}
}

// Dropped determines whether certain events responsible for building the internal state are kept or dropped before hitting
// the output channel.
func (k Ktype) Dropped(capture bool) bool {
	switch k {
	case EnumProcess, EnumThread, FileRundown, FileOpEnd, ReleaseFile, EnumImage, RegCreateKCB, RegKCBRundown:
		if capture {
			return false
		}
		return true
	default:
		return false
	}
}

// Pack transforms event provider GUID and the op code into `Ktype` type. The type provides a convenient way
// to compare different kernel event types.
func Pack(g syscall.GUID, opcode uint8) Ktype {
	return Ktype([17]byte{
		byte(g.Data1 >> 24), byte(g.Data1 >> 16), byte(g.Data1 >> 8), byte(g.Data1),
		byte(g.Data2 >> 8), byte(g.Data2), byte(g.Data3 >> 8), byte(g.Data3),
		g.Data4[0], g.Data4[1], g.Data4[2], g.Data4[3], g.Data4[4], g.Data4[5], g.Data4[6], g.Data4[7],
		byte(opcode),
	})
}
