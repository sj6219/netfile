// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"strings"
	"syscall"
	// "path"
	// "path/filepath"
	// "strings"
	// "syscall"
)

var dll syscall.Handle

type NetFileName struct {
	Server string
	Path   string
}

func LoadLibrary(name string) {
	dll, _ = syscall.LoadLibrary(name)
}

func Syscall(name string, nargs uintptr, a1, a2, a3 uintptr) (r1, r2 uintptr, e syscall.Errno) {
	f, _ := syscall.GetProcAddress(dll, name)
	return syscall.Syscall(uintptr(f), nargs, a1, a2, a3)
}

func (name NetFileName) DebugString() string {
	if len(name.Server) == 0 {
		return string(name.Path)
	}
	if len(name.Path) == 0 {
		return "."
	}
	i := strings.IndexByte(name.Path, '/')
	if i < 0 {
		i = len(name.Path)
	}
	prefix := name.Path[:i]
	j := strings.IndexByte(prefix, '#')
	if j < 0 {
		return "\\\\" + name.Server + "\\" + name.Path
	}
	return "\\\\" + name.Path[:j] + "\\" + name.Path[j+1:]
}

func (name NetFileName) String() string {
	if len(name.Server) == 0 {
		return string(name.Path)
	}
	return "\\\\" + name.Server + "#" + name.Path
}
