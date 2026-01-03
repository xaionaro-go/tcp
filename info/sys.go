// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package info

import "github.com/xaionaro-go/tcp/opt"

func init() {
	for _, o := range options {
		if o.name == 0 || o.parseFn == nil {
			continue
		}
		opt.Register(o.level, o.name, o.parseFn)
	}
}

const (
	ianaProtocolTCP = 0x6
)

const (
	soInfo = iota
	soCCInfo
	soCCAlgo
	soMax
)

// An option represents a binding for socket option.
type option struct {
	level   int // option level
	name    int // option name, must be equal or greater than 1
	parseFn func([]byte) (opt.Option, error)
}
