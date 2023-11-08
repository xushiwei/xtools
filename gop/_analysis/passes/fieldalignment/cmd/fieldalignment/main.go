// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"github.com/goplus/xtools/gop/analysis/passes/fieldalignment"
	"github.com/goplus/xtools/gop/analysis/singlechecker"
)

func main() { singlechecker.Main(fieldalignment.Analyzer) }
