// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_openbsd.go

package syscall

const (
	SysFIONREAD = 0x4004667f

	SysAF_INET  = 0x2
	SysAF_INET6 = 0x18

	SysPF_INOUT = 0x0
	SysPF_IN    = 0x1
	SysPF_OUT   = 0x2
	sysPF_FWD   = 0x3

	SysDIOCNATLOOK = 0xc0504417
)

type sockaddrStorage struct {
	Len        uint8
	Family     uint8
	X__ss_pad1 [6]uint8
	X__ss_pad2 uint64
	X__ss_pad3 [240]uint8
}

type sockaddr struct {
	Len    uint8
	Family uint8
	Data   [14]int8
}

type SockaddrInet struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type SockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type PfiocNatlook struct {
	Saddr     [16]byte /* pf_addr */
	Daddr     [16]byte /* pf_addr */
	Rsaddr    [16]byte /* pf_addr */
	Rdaddr    [16]byte /* pf_addr */
	Rdomain   uint16
	Rrdomain  uint16
	Sport     uint16
	Dport     uint16
	Rsport    uint16
	Rdport    uint16
	Af        uint8
	Proto     uint8
	Direction uint8
	Pad_cgo_0 [1]byte
}

const (
	SizeofSockaddrStorage = 0x100
	SizeofSockaddr        = 0x10
	SizeofSockaddrInet    = 0x10
	SizeofSockaddrInet6   = 0x1c
	SizeofPfiocNatlook    = 0x50
)
