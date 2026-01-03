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
	SoAvailable: {SysSOL_SOCKET, SysSO_NWRITE},
}

func (nl *PfiocNatlook) ReadPort() int {
	return int(binary.BigEndian.Uint16(nl.Rdxport[:2]))
}

func (nl *PfiocNatlook) SetPort(remote, local int) {
	binary.BigEndian.PutUint16((*[2]byte)(unsafe.Pointer(&nl.Sxport))[:2], uint16(remote))
	binary.BigEndian.PutUint16((*[2]byte)(unsafe.Pointer(&nl.Dxport))[:2], uint16(local))
}
