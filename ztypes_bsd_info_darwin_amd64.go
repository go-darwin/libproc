// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -srcdir . -godefs /Users/zchee/go/src/github.com/go-darwin/libproc/bsd_info_darwin_amd64.go

package libproc

const MaxComLen = 0x10

type ProcBsdinfo struct {
	Pbi_flags        uint32
	Pbi_status       uint32
	Pbi_xstatus      uint32
	Pbi_pid          uint32
	Pbi_ppid         uint32
	Pbi_uid          uint32
	Pbi_gid          uint32
	Pbi_ruid         uint32
	Pbi_rgid         uint32
	Pbi_svuid        uint32
	Pbi_svgid        uint32
	Rfu_1            uint32
	Pbi_comm         [16]int8
	Pbi_name         [32]int8
	Pbi_nfiles       uint32
	Pbi_pgid         uint32
	Pbi_pjobc        uint32
	E_tdev           uint32
	E_tpgid          uint32
	Pbi_nice         int32
	Pbi_start_tvsec  uint64
	Pbi_start_tvusec uint64
}
