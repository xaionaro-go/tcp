// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import (
	"encoding/binary"
	"unsafe"
)

var NativeEndian binary.ByteOrder

func init() {
	i := uint32(1)
	b := (*[4]byte)(unsafe.Pointer(&i))
	if b[0] == 1 {
		NativeEndian = binary.LittleEndian
	} else {
		NativeEndian = binary.BigEndian
	}
}

const (
	IANAProtocolIP   = 0x0
	IANAProtocolTCP  = 0x6
	IANAProtocolIPv6 = 0x29
)

const (
	SoBuffered = iota
	SoAvailable
	EndOfSo
)

type Option struct {
	Level int // Option level
	Name  int // Option name, must be equal or greater than 1
}
