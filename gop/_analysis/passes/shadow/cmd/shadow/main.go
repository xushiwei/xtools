// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The shadow command runs the shadow analyzer.
package main

import (
	"github.com/goplus/xtools/gop/analysis/passes/shadow"
	"github.com/goplus/xtools/gop/analysis/singlechecker"
)

func main() { singlechecker.Main(shadow.Analyzer) }
