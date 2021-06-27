// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

package libproc

/*
#cgo CFLAGS: -mmacosx-version-min=12.0

#include <sys/msgbuf.h>
*/
import "C"

const (
	MaxMsgBSize = C.MAX_MSG_BSIZE
)
