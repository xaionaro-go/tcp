// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import (
	"encoding/binary"
	"unsafe"
)

var options = [EndOfSo]Option{
	SoBuffered:  {0, SysFIONREAD},
	SoAvailable: {0, SysFIONSPACE},
}

func (nl *PfiocNatlook) ReadPort() int {
	return int(binary.BigEndian.Uint16((*[2]byte)(unsafe.Pointer(&nl.Rdport))[:]))
}

func (nl *PfiocNatlook) SetPort(remote, local int) {
	binary.BigEndian.PutUint16((*[2]byte)(unsafe.Pointer(&nl.Sport))[:], uint16(remote))
	binary.BigEndian.PutUint16((*[2]byte)(unsafe.Pointer(&nl.Dport))[:], uint16(local))
}
