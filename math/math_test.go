// Copyright 2018 The go-python Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math_test

import (
	"testing"

	"github.com/go-python/gpython/pytest"
)

func TestVm(t *testing.T) {
	pytest.RunTests(t, "tests")
}
