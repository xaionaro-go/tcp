// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_freebsd.go

package syscall

const (
	SysFIONREAD  = 0x4004667f
	SysFIONWRITE = 0x40046677
	SysFIONSPACE = 0x40046676

	SysAF_INET  = 0x2
	SysAF_INET6 = 0x1c

	SysPF_INOUT = 0x0
	SysPF_IN    = 0x1
	SysPF_OUT   = 0x2
	sysPF_FWD   = 0x3

	SysDIOCNATLOOK = 0xc04c4417
)

type sockaddrStorage struct {
	Len         uint8
	Family      uint8
	X__ss_pad1  [6]int8
	X__ss_align int64
	X__ss_pad2  [112]int8
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
	SizeofSockaddrStorage = 0x80
	SizeofSockaddr        = 0x10
	SizeofSockaddrInet    = 0x10
	SizeofSockaddrInet6   = 0x1c
	SizeofPfiocNatlook    = 0x4c
)
