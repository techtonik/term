// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package term

import (
	"os"
	"testing"
)

func TestIsTerminal(t *testing.T) {
	file, err := os.Open("term_test.go")
	if err != nil {
		t.Error(err)
	}

	if IsTerminal(file.Fd()) {
		t.Error("Open file is detected as terminal")
	}

	if !IsTerminal(os.Stdin.Fd()) {
		t.Error("Stdin is not terminal")
	}

	file.Close()
}
