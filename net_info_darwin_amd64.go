// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

// for in4in6_addr
// +godefs map struct_in_addr  [4]byte // in_addr

// for in_sockinfo
// +godefs map struct___1 InsiV4 // insi_v4
// +godefs map struct___2 InsiV6 // insi_v6

package libproc

/*
#cgo CFLAGS: -mmacosx-version-min=12.0

#include <stdlib.h>

#include <sys/proc_info.h>

struct insi_v4 {
	u_char in4_tos;
};

struct insi_v6 {
	uint8_t                 in6_hlim;
	int                     in6_cksum;
	u_short                 in6_ifindex;
	short                   in6_hops;
};

// local host table entry
union insi_laddr {
	struct in4in6_addr      ina_46;
	struct in6_addr         ina_6;
};
*/
import "C"

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L614-L617
type SocketFDinfo C.struct_socket_fdinfo

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L298-L304
type ProcFileinfo C.struct_proc_fileinfo

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L714-L723
const (
	SockinfoGeneric   = C.SOCKINFO_GENERIC
	SockinfoIn        = C.SOCKINFO_IN
	SockinfoKernCtl   = C.SOCKINFO_KERN_CTL
	SockinfoKernEvent = C.SOCKINFO_KERN_EVENT
	SockinfoNdrv      = C.SOCKINFO_NDRV
	SockinfoTCP       = C.SOCKINFO_TCP
	SockinfoUn        = C.SOCKINFO_UN
	SockinfoVsock     = C.SOCKINFO_VSOCK
)

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L583-L612
type SocketInfo C.struct_socket_info

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L332-L357
type VInfoStat C.struct_vinfo_stat

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L562-L570
type SockbufInfo C.struct_sockbuf_info

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/socket.h#L612-L619
type SockProto C.struct_sockproto

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L414-L417
type In4In6Addr C.struct_in4in6_addr

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L419-L447
type InSockinfo C.struct_in_sockinfo

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L438-L440
type InsiV4 C.struct_insi_v4

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L441-L446
type InsiV6 C.struct_insi_v6

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L434-L437
type InsiLaddr C.union_insi_laddr

type TCPSIState C.int

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L459-L470
const (
	TSI_S_CLOSED       TCPSIState = C.TSI_S_CLOSED       // closed
	TSI_S_LISTEN       TCPSIState = C.TSI_S_LISTEN       // listening for connection
	TSI_S_SYN_SENT     TCPSIState = C.TSI_S_SYN_SENT     // active, have sent syn
	TSI_S_SYN_RECEIVED TCPSIState = C.TSI_S_SYN_RECEIVED // have send and received syn
	TSI_S_ESTABLISHED  TCPSIState = C.TSI_S_ESTABLISHED  // established
	TSI_S__CLOSE_WAIT  TCPSIState = C.TSI_S__CLOSE_WAIT  // rcvd fin, waiting for close
	TSI_S_FIN_WAIT_1   TCPSIState = C.TSI_S_FIN_WAIT_1   // have closed, sent fin
	TSI_S_CLOSING      TCPSIState = C.TSI_S_CLOSING      // closed xchd FIN; await FIN ACK
	TSI_S_LAST_ACK     TCPSIState = C.TSI_S_LAST_ACK     // had fin and close; await FIN ACK
	TSI_S_FIN_WAIT_2   TCPSIState = C.TSI_S_FIN_WAIT_2   // have closed, fin is acked
	TSI_S_TIME_WAIT    TCPSIState = C.TSI_S_TIME_WAIT    // in 2*msl quiet wait after close
	TSI_S_RESERVED     TCPSIState = C.TSI_S_RESERVED     // pseudo state: reserved
)

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L472-L480
type TCPSockinfo C.struct_tcp_sockinfo

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L487-L498
type UnSocinfo C.struct_un_sockinfo

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L504-L508
type NdrvInfo C.struct_ndrv_info

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L514-L518
type KernEventInfo C.struct_kern_event_info

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L524-L532
type KernCTLInfo C.struct_kern_ctl_info

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L538-L543
type VsockSockinfo C.struct_vsock_sockinfo
