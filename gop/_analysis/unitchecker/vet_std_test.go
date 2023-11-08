// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unitchecker_test

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

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
	"github.com/goplus/xtools/gop/analysis/passes/unusedresult"
	"github.com/goplus/xtools/gop/analysis/unitchecker"
)

// vet is the entrypoint of this executable when ENTRYPOINT=vet.
// Keep consistent with the actual vet in GOROOT/src/cmd/vet/main.go.
func vet() {
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
		// unsafeptr.Analyzer, // currently reports findings in runtime
		unusedresult.Analyzer,
	)
}

// TestVetStdlib runs the same analyzers as the actual vet over the
// standard library, using go vet and unitchecker, to ensure that
// there are no findings.
func TestVetStdlib(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in -short mode")
	}
	if version := runtime.Version(); !strings.HasPrefix(version, "devel") {
		t.Skipf("This test is only wanted on development branches where code can be easily fixed. Skipping because runtime.Version=%q.", version)
	}

	cmd := exec.Command("go", "vet", "-vettool="+os.Args[0], "std")
	cmd.Env = append(os.Environ(), "ENTRYPOINT=vet")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Errorf("go vet std failed (%v):\n%s", err, out)
	}
}
