// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fieldalignment_test

import (
	"testing"

	"github.com/goplus/xtools/gop/analysis/analysistest"
	"github.com/goplus/xtools/gop/analysis/passes/fieldalignment"
)

func TestTest(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, fieldalignment.Analyzer, "a")
}
