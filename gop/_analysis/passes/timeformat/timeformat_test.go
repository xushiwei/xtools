// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timeformat_test

import (
	"testing"

	"github.com/goplus/xtools/gop/analysis/analysistest"
	"github.com/goplus/xtools/gop/analysis/passes/timeformat"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, timeformat.Analyzer, "a")
}
