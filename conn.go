// Copyright 2014 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tcp

import (
	"net"
	"syscall"
	"time"

	"github.com/xaionaro-go/tcp/opt"
)

var _ net.Conn = &Conn{}

// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Read reads data from the connection.
func (c *Conn) Read(b []byte) (int, error) {
	return c.Conn.Read(b)
}

// Write writes data to the connection.
func (c *Conn) Write(b []byte) (int, error) {
	return c.Conn.Write(b)
}

// SetDeadline sets the read and write deadlines associated
func (c *Conn) SetDeadline(t time.Time) error {
	return c.Conn.SetDeadline(t)
}

// SetReadDeadline sets the deadline for future Read calls
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.Conn.SetReadDeadline(t)
}

// SetWriteDeadline sets the deadline for future Write calls
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.Conn.SetWriteDeadline(t)
}

// Close closes the connection.
func (c *Conn) Close() error {
	return c.Conn.Close()
}

// SetOption sets a socket option.
func (c *Conn) SetOption(o opt.Option) error {
	if !c.ok() {
		return syscall.EINVAL
	}
	b, err := o.Marshal()
	if err != nil {
		return &net.OpError{Op: "raw-control", Net: c.LocalAddr().Network(), Source: nil, Addr: c.LocalAddr(), Err: err}
	}
	if err := c.setOption(o.Level(), o.Name(), b); err != nil {
		return &net.OpError{Op: "raw-control", Net: c.LocalAddr().Network(), Source: nil, Addr: c.LocalAddr(), Err: err}
	}
	return nil
}

// Option returns a socket option.
func (c *Conn) Option(level, name int, b []byte) (opt.Option, error) {
	if !c.ok() || len(b) == 0 {
		return nil, syscall.EINVAL
	}
	n, err := c.option(level, name, b)
	if err != nil {
		return nil, &net.OpError{Op: "raw-control", Net: c.LocalAddr().Network(), Source: nil, Addr: c.LocalAddr(), Err: err}
	}
	o, err := opt.Parse(level, name, b[:n])
	if err != nil {
		return nil, &net.OpError{Op: "raw-control", Net: c.LocalAddr().Network(), Source: nil, Addr: c.LocalAddr(), Err: err}
	}
	return o, nil
}

// Buffered returns the number of bytes that can be read from the
// underlying socket read buffer.
// It returns -1 when the platform doesn't support this feature.
func (c *Conn) Buffered() int {
	if !c.ok() {
		return -1
	}
	return c.buffered()
}

// Available returns how many bytes are unused in the underlying
// socket write buffer.
// It returns -1 when the platform doesn't support this feature.
func (c *Conn) Available() int {
	if !c.ok() {
		return -1
	}
	return c.available()
}

// OriginalDst returns an original destination address, which is an
// address not modified by intermediate entities such as network
// address and port translators inside the kernel, on the connection.
//
// Only Linux and BSD variants using PF support this feature.
func (c *Conn) OriginalDst() (net.Addr, error) {
	if !c.ok() {
		return nil, syscall.EINVAL
	}
	la := c.LocalAddr().(*net.TCPAddr)
	od, err := c.originalDst(la, c.RemoteAddr().(*net.TCPAddr))
	if err != nil {
		return nil, &net.OpError{Op: "raw-control", Net: c.LocalAddr().Network(), Source: nil, Addr: la, Err: err}
	}
	return od, nil
}
