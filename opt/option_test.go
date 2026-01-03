// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opt_test

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/xaionaro-go/tcp/opt"
)

func TestMarshalAndParse(t *testing.T) {
	opts := []opt.Option{
		opt.NoDelay(true),
		opt.SendBuffer(1<<16 - 1),
		opt.ReceiveBuffer(1<<16 - 1),
		opt.KeepAlive(true),
		opt.Linger{OnOff: true, Linger: 10 * time.Second},
	}
	switch runtime.GOOS {
	case "openbsd":
	case "windows":
		opts = append(opts, opt.KeepAliveIdleInterval(1*time.Hour))
		opts = append(opts, opt.KeepAliveProbeInterval(10*time.Minute))
	default:
		opts = append(opts, opt.KeepAliveIdleInterval(1*time.Hour))
		opts = append(opts, opt.KeepAliveProbeInterval(10*time.Minute))
		opts = append(opts, opt.KeepAliveProbeCount(3))
	}
	switch runtime.GOOS {
	case "netbsd", "windows":
	default:
		opts = append(opts, opt.Cork(true))
	}
	switch runtime.GOOS {
	case "darwin", "linux":
		opts = append(opts, opt.NotSentLowWMK(1))
	}
	switch runtime.GOOS {
	case "linux":
		opts = append(opts, opt.QuickAck(true))
		opts = append(opts, opt.ThinDupAck(true))
		opts = append(opts, opt.ThinLinearTimeouts(true))
	}
	switch runtime.GOOS {
	case "windows":
	default:
		opts = append(opts, opt.MSS(4092))
		opts = append(opts, opt.Error(42))
	}
	switch runtime.GOOS {
	case "darwin":
		opts = append(opts, opt.ECN(true))
	}

	for _, o := range opts {
		if o.Level() <= 0 {
			t.Fatalf("got %#x; want greater than zero", o.Level())
		}
		if o.Name() <= 0 {
			t.Fatalf("got %#x; want greater than zero", o.Name())
		}
		b, err := o.Marshal()
		if err != nil {
			t.Fatal(err)
		}
		if runtime.GOOS == "windows" {
			continue
		}
		oo, err := opt.Parse(o.Level(), o.Name(), b)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(oo, o) {
			t.Fatalf("got %#v; want %#v", oo, o)
		}
	}
}

const (
	testOptLevel = 0xfff1
	testOptName  = 0xfff2
)

type testOption struct{}

func (*testOption) Level() int                     { return testOptLevel }
func (*testOption) Name() int                      { return testOptName }
func (*testOption) Marshal() ([]byte, error)       { return make([]byte, 16), nil }
func parseTestOption(_ []byte) (opt.Option, error) { return &testOption{}, nil }

func TestUserDefinedOptionParser(t *testing.T) {
	var b [16]byte
	opt.Register(testOptLevel, testOptName, parseTestOption)
	o, err := opt.Parse(testOptLevel, testOptName, b[:])
	if err != nil {
		t.Fatal(err)
	}
	opt.Unregister(testOptLevel, testOptName)
	o, err = opt.Parse(testOptLevel, testOptName, b[:])
	if err == nil || o != nil {
		t.Fatalf("got %v, %v; want nil, error", o, err)
	}
}

func TestParseWithVariousBufferLengths(t *testing.T) {
	for _, o := range []opt.Option{
		opt.NoDelay(true),
		opt.SendBuffer(1<<16 - 1),
		opt.ReceiveBuffer(1<<16 - 1),
		opt.KeepAlive(true),
		opt.KeepAliveIdleInterval(1 * time.Hour),
		opt.KeepAliveProbeInterval(10 * time.Minute),
		opt.KeepAliveProbeCount(3),
		opt.Error(42),
	} {
		for i := 0; i < 256; i++ {
			b := make([]byte, i)
			if _, err := opt.Parse(o.Level(), o.Name(), b); err == nil {
				break
			}
		}
	}
}
