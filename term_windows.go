// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package term

// Returns true if the given file descriptor points to terminal or console.
func IsTerminal(fd uintptr) bool {
	return false
}
