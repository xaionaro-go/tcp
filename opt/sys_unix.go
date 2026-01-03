// Copyright 2016 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package opt

import (
	"errors"
	"time"
	"unsafe"
)

// Marshal implements the Marshal method of Option interface.
func (nd NoDelay) Marshal() ([]byte, error) {
	v := boolint32(bool(nd))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (mss MSS) Marshal() ([]byte, error) {
	v := int32(mss)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (sb SendBuffer) Marshal() ([]byte, error) {
	v := int32(sb)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (rb ReceiveBuffer) Marshal() ([]byte, error) {
	v := int32(rb)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAlive) Marshal() ([]byte, error) {
	v := boolint32(bool(ka))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAliveIdleInterval) Marshal() ([]byte, error) {
	ka += KeepAliveIdleInterval(options[soKeepidle].uot - time.Nanosecond)
	v := int32(time.Duration(ka) / options[soKeepidle].uot)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAliveProbeInterval) Marshal() ([]byte, error) {
	ka += KeepAliveProbeInterval(options[soKeepintvl].uot - time.Nanosecond)
	v := int32(time.Duration(ka) / options[soKeepintvl].uot)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ka KeepAliveProbeCount) Marshal() ([]byte, error) {
	v := int32(ka)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ck Cork) Marshal() ([]byte, error) {
	v := boolint32(bool(ck))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (ns NotSentLowWMK) Marshal() ([]byte, error) {
	v := int32(ns)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (e Error) Marshal() ([]byte, error) {
	v := int32(e)
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (l Linger) Marshal() ([]byte, error) {
	v := struct {
		OnOff  int32
		Linger int32
	}{
		OnOff:  boolint32(l.OnOff),
		Linger: int32(l.Linger / time.Second),
	}
	return (*[8]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (cn ECN) Marshal() ([]byte, error) {
	v := boolint32(bool(cn))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (qa QuickAck) Marshal() ([]byte, error) {
	v := boolint32(bool(qa))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (tda ThinDupAck) Marshal() ([]byte, error) {
	v := boolint32(bool(tda))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

// Marshal implements the Marshal method of Option interface.
func (tlt ThinLinearTimeouts) Marshal() ([]byte, error) {
	v := boolint32(bool(tlt))
	return (*[4]byte)(unsafe.Pointer(&v))[:], nil
}

func (t TOA) Marshal() ([]byte, error) {
	opt := []byte{t.Kind, t.Len}
	bp := (*[2]byte)(unsafe.Pointer(&t.Port))[:]
	data := append(opt, bp...)
	data = append(data, t.Ip...)
	return data, nil
}

func parseNoDelay(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return NoDelay(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseMSS(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return MSS(nativeEndian.Uint32(b)), nil
}

func parseSendBuffer(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return SendBuffer(nativeEndian.Uint32(b)), nil
}

func parseReceiveBuffer(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return ReceiveBuffer(nativeEndian.Uint32(b)), nil
}

func parseKeepAlive(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return KeepAlive(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseKeepAliveIdleInterval(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	v := time.Duration(nativeEndian.Uint32(b)) * options[soKeepidle].uot
	return KeepAliveIdleInterval(v), nil
}

func parseKeepAliveProbeInterval(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	v := time.Duration(nativeEndian.Uint32(b)) * options[soKeepintvl].uot
	return KeepAliveProbeInterval(v), nil
}

func parseKeepAliveProbeCount(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return KeepAliveProbeCount(nativeEndian.Uint32(b)), nil
}

func parseCork(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return Cork(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseNotSentLowWMK(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return NotSentLowWMK(nativeEndian.Uint32(b)), nil
}

func parseError(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return Error(nativeEndian.Uint32(b)), nil
}

func parseLinger(b []byte) (Option, error) {
	if len(b) < 8 {
		return nil, errors.New("short buffer")
	}
	v := (*struct {
		OnOff  int32
		Linger int32
	})(unsafe.Pointer(&b[0]))
	return Linger{
		OnOff:  uint32bool(uint32(v.OnOff)),
		Linger: time.Duration(v.Linger) * time.Second,
	}, nil
}

func parseECN(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return ECN(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseQuickAck(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return QuickAck(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseThinDupAck(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return ThinDupAck(uint32bool(nativeEndian.Uint32(b))), nil
}

func parseThinLinearTimeouts(b []byte) (Option, error) {
	if len(b) < 4 {
		return nil, errors.New("short buffer")
	}
	return ThinLinearTimeouts(uint32bool(nativeEndian.Uint32(b))), nil
}
