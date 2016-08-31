// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package obj

import "os"

const (
	ABase386 = (1 + iota) << 12
	ABaseARM
	ABaseAMD64
	ABasePPC64
	ABaseARM64
	ABaseMIPS64
	AMask = 1<<12 - 1 // AND with this to use the opcode as an array index.
)

const (
	// Because of masking operations in the encodings, each register
	// space should start at 0 modulo some power of 2.
	RBase386   = 1 * 1024
	RBaseAMD64 = 2 * 1024
	RBaseARM   = 3 * 1024
	RBasePPC64 = 4 * 1024 // range [4k, 8k)
	RBaseARM64 = 8 * 1024 // range [8k, 12k)
	RBaseMIPS64 = 13 * 1024 // range [13k, 14k)
)

func envOr(key, value string) string {
	if x := os.Getenv(key); x != "" {
		return x
	}
	return value
}

func Getgoos() string {
	return envOr("GOOS", "")
}
