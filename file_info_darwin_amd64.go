// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

package libproc

/*
#cgo CFLAGS: -mmacosx-version-min=12.0

#include <sys/proc_info.h>
*/
import "C"

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L727
type ProcFDType C.uint32_t

// defns of process file desc type
// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L714-L723
const (
	ProxFDTypeAtalk     ProcFDType = C.PROX_FDTYPE_ATALK
	ProxFDTypeVnode     ProcFDType = C.PROX_FDTYPE_VNODE
	ProxFDTypeSocket    ProcFDType = C.PROX_FDTYPE_SOCKET
	ProxFDTypePshm      ProcFDType = C.PROX_FDTYPE_PSHM
	ProxFDTypePsem      ProcFDType = C.PROX_FDTYPE_PSEM
	ProxFDTypeKqueue    ProcFDType = C.PROX_FDTYPE_KQUEUE
	ProxFDTypePipe      ProcFDType = C.PROX_FDTYPE_PIPE
	ProxFDTypeFsevents  ProcFDType = C.PROX_FDTYPE_FSEVENTS
	ProxFDTypeNetpolicy ProcFDType = C.PROX_FDTYPE_NETPOLICY
)
