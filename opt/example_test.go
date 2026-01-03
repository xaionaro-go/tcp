// Copyright 2017 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opt_test

import (
	"log"
	"net"
	"testing"
	"time"

	"github.com/xaionaro-go/tcp"
	"github.com/xaionaro-go/tcp/opt"
)

func TestOption(t *testing.T) {
	c, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	tc, err := tcp.NewConn(c)
	if err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(opt.KeepAlive(true)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(opt.MSS(1100)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(opt.TOA{
		Port: 9000,
		Ip:   net.ParseIP("127.0.0.5"),
	}); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(opt.KeepAliveIdleInterval(3 * time.Minute)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(opt.KeepAliveProbeInterval(30 * time.Second)); err != nil {
		log.Fatal(err)
	}
	if err := tc.SetOption(opt.KeepAliveProbeCount(3)); err != nil {
		log.Fatal(err)
	}

	if err := c.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
		log.Fatal(err)
	}

	log.Println("hello world!")

}
