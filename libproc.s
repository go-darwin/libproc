// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

#include "textflag.h"

// int proc_listallpids(void * buffer, int buffersize)
GLOBL ·libc_proc_listallpids_trampoline_addr(SB), RODATA, $8
DATA ·libc_proc_listallpids_trampoline_addr(SB)/8, $libc_proc_listallpids_trampoline<>(SB)

TEXT libc_proc_listallpids_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_proc_listallpids(SB)

// int proc_pidpath(int pid, void * buffer, uint32_t buffersize)
GLOBL ·libc_proc_pidpath_trampoline_addr(SB), RODATA, $8
DATA ·libc_proc_pidpath_trampoline_addr(SB)/8, $libc_proc_pidpath_trampoline<>(SB)

TEXT libc_proc_pidpath_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_proc_pidpath(SB)
