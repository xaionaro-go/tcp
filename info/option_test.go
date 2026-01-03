// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package info_test

import (
	"runtime"
	"testing"

	"github.com/xaionaro-go/tcp/info"
	"github.com/xaionaro-go/tcp/opt"
)

func TestMarshalAndParse(t *testing.T) {
	opts := make([]opt.Option, 0, 3)
	switch runtime.GOOS {
	case "darwin", "freebsd", "netbsd":
		opts = append(opts, &info.Info{})
	case "linux":
		opts = append(opts, &info.Info{})
		opts = append(opts, &info.CCInfo{})
		opts = append(opts, info.CCAlgorithm("vegas"))
	default:
		t.Skipf("%s/%s", runtime.GOOS, runtime.GOARCH)
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
		if len(b) == 0 {
			continue
		}
		oo, err := opt.Parse(o.Level(), o.Name(), b)
		if err != nil {
			t.Fatal(err)
		}
		if oo, ok := oo.(*info.Info); ok {
			if _, err := oo.MarshalJSON(); err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestParseWithVariousBufferLengths(t *testing.T) {
	for _, o := range []opt.Option{
		&info.Info{},
		&info.CCInfo{},
		info.CCAlgorithm("vegas"),
	} {
		for i := 0; i < 256; i++ {
			b := make([]byte, i)
			if _, err := opt.Parse(o.Level(), o.Name(), b); err == nil {
				break
			}
		}
	}
}
