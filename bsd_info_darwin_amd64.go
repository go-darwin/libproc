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

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/param.h#L95
const MaxComLen = C.MAXCOMLEN

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L63-L86
type ProcBsdinfo C.struct_proc_bsdinfo
