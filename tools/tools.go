// SPDX-FileCopyrightText: 2021 The go-darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build tools
// +build tools

package tools

import (
	_ "github.com/klauspost/asmfmt/cmd/asmfmt"
	_ "go-darwin.dev/tools/cmd/asmvet"
)
