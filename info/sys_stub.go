// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !darwin && !freebsd && !linux && !netbsd
// +build !darwin,!freebsd,!linux,!netbsd

package info

import (
	"errors"

	"github.com/xaionaro-go/tcp/opt"
)

var options [soMax]option

// Marshal implements the Marshal method of opt.Option interface.
func (i *Info) Marshal() ([]byte, error) {
	return nil, errors.New("operation not supported")
}

// A SysInfo represents platform-specific information.
type SysInfo struct{}

func parseInfo(b []byte) (opt.Option, error) {
	return nil, errors.New("operation not supported")
}

func parseCCAlgorithmInfo(name string, b []byte) (CCAlgorithmInfo, error) {
	return nil, errors.New("operation not supported")
}
