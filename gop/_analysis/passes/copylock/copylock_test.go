// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package copylock_test

import (
	"testing"

	"github.com/goplus/xtools/gop/analysis/analysistest"
	"github.com/goplus/xtools/gop/analysis/passes/copylock"
	"golang.org/x/tools/internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	pkgs := []string{"a"}
	if typeparams.Enabled {
		pkgs = append(pkgs, "typeparams")
	}
	analysistest.Run(t, testdata, copylock.Analyzer, pkgs...)
}
