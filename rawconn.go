// Copyright 2017 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

import (
	"errors"
	"net"
	"os"
	"runtime"
	"syscall"

	"github.com/xaionaro-go/tcp/opt"
	xsyscall "github.com/xaionaro-go/tcp/syscall"
)

// A Conn represents an end point that uses TCP connection.
// It allows to set non-portable, platform-dependent TCP-level socket
// options.
type Conn struct {
	Conn    net.Conn
	RawConn syscall.RawConn
}

func (c *Conn) ok() bool { return c != nil && c.Conn != nil && c.RawConn != nil }

func (c *Conn) setOption(level, name int, b []byte) error {
	var operr error
	fn := func(s uintptr) {
		operr = xsyscall.Setsockopt(s, level, name, b)
	}
	if err := c.RawConn.Control(fn); err != nil {
		return err
	}
	return os.NewSyscallError("Setsockopt", operr)
}

func (c *Conn) option(level, name int, b []byte) (int, error) {
	var operr error
	var n int
	fn := func(s uintptr) {
		n, operr = xsyscall.Getsockopt(s, level, name, b)
	}
	if err := c.RawConn.Control(fn); err != nil {
		return 0, err
	}
	return n, os.NewSyscallError("Getsockopt", operr)
}

func (c *Conn) buffered() int {
	var operr error
	var n int
	fn := func(s uintptr) {
		var b [4]byte
		operr = xsyscall.Ioctl(s, xsyscall.OptionByKey(xsyscall.SoBuffered).Name, b[:])
		if operr != nil {
			return
		}
		n = int(xsyscall.NativeEndian.Uint32(b[:]))
	}
	err := c.RawConn.Control(fn)
	if err != nil || operr != nil {
		return -1
	}
	return n
}

func (c *Conn) available() int {
	var operr error
	var n int
	fn := func(s uintptr) {
		var b [4]byte
		if runtime.GOOS == "darwin" {
			_, operr = xsyscall.Getsockopt(s, xsyscall.OptionByKey(xsyscall.SoAvailable).Level, xsyscall.OptionByKey(xsyscall.SoAvailable).Name, b[:])
		} else {
			operr = xsyscall.Ioctl(s, xsyscall.OptionByKey(xsyscall.SoAvailable).Name, b[:])
		}
		if operr != nil {
			return
		}
		n = int(xsyscall.NativeEndian.Uint32(b[:]))
		if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
			var o opt.SendBuffer
			_, operr = xsyscall.Getsockopt(s, o.Level(), o.Name(), b[:])
			if operr != nil {
				return
			}
			n = int(xsyscall.NativeEndian.Uint32(b[:])) - n
		}
	}
	err := c.RawConn.Control(fn)
	if err != nil || operr != nil {
		return -1
	}
	return n
}

// NewConn returns a new end point.
func NewConn(c net.Conn) (*Conn, error) {
	type tcpConn interface {
		SyscallConn() (syscall.RawConn, error)
		SetLinger(int) error
	}
	var _ tcpConn = &net.TCPConn{}
	cc := &Conn{Conn: c}
	switch c := c.(type) {
	case tcpConn:
		var err error
		cc.RawConn, err = c.SyscallConn()
		if err != nil {
			return nil, err
		}
		return cc, nil
	default:
		return nil, errors.New("unknown connection type")
	}
}
