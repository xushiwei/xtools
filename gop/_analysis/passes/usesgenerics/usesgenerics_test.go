// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package usesgenerics_test

import (
	"testing"

	"github.com/goplus/xtools/gop/analysis/analysistest"
	"github.com/goplus/xtools/gop/analysis/passes/usesgenerics"
	"golang.org/x/tools/internal/typeparams"
)

func Test(t *testing.T) {
	if !typeparams.Enabled {
		t.Skip("type parameters are not enabled at this Go version")
	}
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, usesgenerics.Analyzer, "a", "b", "c", "d")
}
