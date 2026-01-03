// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

import (
	"encoding/binary"
	"net"
	"unsafe"

	xsyscall "github.com/xaionaro-go/tcp/syscall"
)

func (c *Conn) originalDst(la, _ *net.TCPAddr) (net.Addr, error) {
	var level, name int
	var b []byte
	if la.IP.To4() != nil {
		level = xsyscall.IANAProtocolIP
		name = xsyscall.SysSO_ORIGINAL_DST
		b = make([]byte, xsyscall.SizeofSockaddrInet)
	}
	if la.IP.To16() != nil && la.IP.To4() == nil {
		level = xsyscall.IANAProtocolIPv6
		name = xsyscall.SysIP6T_SO_ORIGINAL_DST
		b = make([]byte, xsyscall.SizeofSockaddrInet6)
	}
	if _, err := c.option(level, name, b); err != nil {
		return nil, err
	}
	od := new(net.TCPAddr)
	switch len(b) {
	case xsyscall.SizeofSockaddrInet:
		sa := (*xsyscall.SockaddrInet)(unsafe.Pointer(&b[0]))
		od.IP = make(net.IP, net.IPv4len)
		copy(od.IP, sa.Addr[:])
		od.Port = int(binary.BigEndian.Uint16((*[2]byte)(unsafe.Pointer(&sa.Port))[:]))
	case xsyscall.SizeofSockaddrInet6:
		sa := (*xsyscall.SockaddrInet6)(unsafe.Pointer(&b[0]))
		od.IP = make(net.IP, net.IPv6len)
		copy(od.IP, sa.Addr[:])
		od.Port = int(binary.BigEndian.Uint16((*[2]byte)(unsafe.Pointer(&sa.Port))[:]))
		od.Zone = zoneCache.name(int(sa.Scope_id))
	}
	return od, nil
}
