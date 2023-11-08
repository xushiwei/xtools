// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This file provides an example command for static checkers
// conforming to the github.com/goplus/xtools/gop/analysis API.
// It serves as a model for the behavior of the cmd/vet tool in $GOROOT.
// Being based on the unitchecker driver, it must be run by go vet:
//
//	$ go build -o unitchecker main.go
//	$ go vet -vettool=unitchecker my/project/...
//
// For a checker also capable of running standalone, use multichecker.
package main

import (
	"github.com/goplus/xtools/gop/analysis/unitchecker"

	"github.com/goplus/xtools/gop/analysis/passes/asmdecl"
	"github.com/goplus/xtools/gop/analysis/passes/assign"
	"github.com/goplus/xtools/gop/analysis/passes/atomic"
	"github.com/goplus/xtools/gop/analysis/passes/bools"
	"github.com/goplus/xtools/gop/analysis/passes/buildtag"
	"github.com/goplus/xtools/gop/analysis/passes/cgocall"
	"github.com/goplus/xtools/gop/analysis/passes/composite"
	"github.com/goplus/xtools/gop/analysis/passes/copylock"
	"github.com/goplus/xtools/gop/analysis/passes/directive"
	"github.com/goplus/xtools/gop/analysis/passes/errorsas"
	"github.com/goplus/xtools/gop/analysis/passes/framepointer"
	"github.com/goplus/xtools/gop/analysis/passes/httpresponse"
	"github.com/goplus/xtools/gop/analysis/passes/ifaceassert"
	"github.com/goplus/xtools/gop/analysis/passes/loopclosure"
	"github.com/goplus/xtools/gop/analysis/passes/lostcancel"
	"github.com/goplus/xtools/gop/analysis/passes/nilfunc"
	"github.com/goplus/xtools/gop/analysis/passes/printf"
	"github.com/goplus/xtools/gop/analysis/passes/shift"
	"github.com/goplus/xtools/gop/analysis/passes/sigchanyzer"
	"github.com/goplus/xtools/gop/analysis/passes/stdmethods"
	"github.com/goplus/xtools/gop/analysis/passes/stringintconv"
	"github.com/goplus/xtools/gop/analysis/passes/structtag"
	"github.com/goplus/xtools/gop/analysis/passes/testinggoroutine"
	"github.com/goplus/xtools/gop/analysis/passes/tests"
	"github.com/goplus/xtools/gop/analysis/passes/timeformat"
	"github.com/goplus/xtools/gop/analysis/passes/unmarshal"
	"github.com/goplus/xtools/gop/analysis/passes/unreachable"
	"github.com/goplus/xtools/gop/analysis/passes/unsafeptr"
	"github.com/goplus/xtools/gop/analysis/passes/unusedresult"
)

func main() {
	unitchecker.Main(
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		directive.Analyzer,
		errorsas.Analyzer,
		framepointer.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		sigchanyzer.Analyzer,
		stdmethods.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		tests.Analyzer,
		testinggoroutine.Analyzer,
		timeformat.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		unsafeptr.Analyzer,
		unusedresult.Analyzer,
	)
}
