// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package opt

// +godefs map struct_in_addr [4]byte /* in_addr */
// +godefs map struct_in6_addr [16]byte /* in6_addr */

/*
#include <sys/ioctl.h>
#include <sys/socket.h>

#include <linux/in.h>
#include <linux/in6.h>
#include <linux/tcp.h>
*/
import "C"

const (
	sysSOL_SOCKET = C.SOL_SOCKET

	sysSO_KEEPALIVE = C.SO_KEEPALIVE
	sysSO_SNDBUF    = C.SO_SNDBUF
	sysSO_RCVBUF    = C.SO_RCVBUF
	sysSO_ERROR     = C.SO_ERROR
	sysSO_LINGER    = C.SO_LINGER

	sysTCP_NODELAY              = C.TCP_NODELAY
	sysTCP_MAXSEG               = C.TCP_MAXSEG
	sysTCP_KEEPIDLE             = C.TCP_KEEPIDLE
	sysTCP_KEEPINTVL            = C.TCP_KEEPINTVL
	sysTCP_KEEPCNT              = C.TCP_KEEPCNT
	sysTCP_CORK                 = C.TCP_CORK
	sysTCP_NOTSENT_LOWAT        = C.TCP_NOTSENT_LOWAT
	sysTCP_QUICKACK             = C.TCP_QUICKACK
	sysTCP_THIN_LINEAR_TIMEOUTS = C.TCP_THIN_LINEAR_TIMEOUTS
	sysTCP_THIN_DUPACK          = C.TCP_THIN_DUPACK
)
