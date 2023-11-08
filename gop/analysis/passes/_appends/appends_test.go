// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appends_test

import (
	"testing"

	"github.com/goplus/xtools/gop/analysis/analysistest"
	"github.com/goplus/xtools/gop/analysis/passes/appends"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	tests := []string{"a", "b"}
	analysistest.Run(t, testdata, appends.Analyzer, tests...)
}
