// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ignore
// +build ignore

package libproc

/*
#cgo CFLAGS: -mmacosx-version-min=12.0

#include <stdbool.h>
#include <stdint.h>

#include <libproc.h>
*/
import "C"

const (
	ProcPidpathinfoSize    = C.PROC_PIDPATHINFO_SIZE
	ProcPidpathinfoMaxsize = C.PROC_PIDPATHINFO_MAXSIZE
)

type PID C.pid_t

// https://github.com/apple-opensource/xnu/blob/7195.101.1/libsyscall/wrappers/libproc/libproc.h#L41-L61
const (
	/*!
	 *       @define PROC_LISTPIDSPATH_PATH_IS_VOLUME
	 *       @discussion This flag indicates that all processes that hold open
	 *               file references on the volume associated with the specified
	 *               path should be returned.
	 */
	ProcListpidspathPathIsVolume = C.PROC_LISTPIDSPATH_PATH_IS_VOLUME

	/*!
	 *       @define PROC_LISTPIDSPATH_EXCLUDE_EVTONLY
	 *       @discussion This flag indicates that file references that were opened
	 *               with the O_EVTONLY flag should be excluded from the matching
	 *               criteria.
	 */
	ProcListpidspathExcludeEvtonly = C.PROC_LISTPIDSPATH_EXCLUDE_EVTONLY
)

// ProcType used to specify what type of processes you are interested
// in in other calls, such as `listpids`.
type ProcType uint32

// https://github.com/apple-opensource/xnu/blob/7195.101.1/bsd/sys/proc_info.h#L55-L61
const (
	ProcAllPids  ProcType = C.PROC_ALL_PIDS
	ProcPgrpOnly ProcType = C.PROC_PGRP_ONLY
	ProcTtyOnly  ProcType = C.PROC_TTY_ONLY
	ProcUidOnly  ProcType = C.PROC_UID_ONLY
	ProcRUIDOnly ProcType = C.PROC_RUID_ONLY
	ProcPPIDOnly ProcType = C.PROC_PPID_ONLY
	ProcKDBGOnly ProcType = C.PROC_KDBG_ONLY
)

type ProcTaskInfo C.struct_proc_taskinfo
