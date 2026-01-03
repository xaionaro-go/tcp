// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_darwin.go

package syscall

const (
	SysSOL_SOCKET = 0xffff

	SysFIONREAD = 0x4004667f

	SysSO_NREAD     = 0x1020
	SysSO_NWRITE    = 0x1024
	SysSO_NUMRCVPKT = 0x1112

	SysAF_INET  = 0x2
	SysAF_INET6 = 0x1e

	SysPF_INOUT = 0
	SysPF_IN    = 1
	SysPF_OUT   = 2

	SysDIOCNATLOOK = 0xc0544417
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
	Sxport    [4]byte
	Dxport    [4]byte
	Rsxport   [4]byte
	Rdxport   [4]byte
	Af        uint8
	Proto     uint8
	Variant   uint8
	Direction uint8
}

const (
	SizeofSockaddrStorage = 0x80
	SizeofSockaddr        = 0x10
	SizeofSockaddrInet    = 0x10
	SizeofSockaddrInet6   = 0x1c
	SizeofPfiocNatlook    = 0x54
)
