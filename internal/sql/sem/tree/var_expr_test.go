// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tree_test

import (
	"testing"

	"github.com/piotrszyma/sqlfmt/internal/sql/parser"
	"github.com/piotrszyma/sqlfmt/internal/sql/sem/tree"
	"github.com/piotrszyma/sqlfmt/internal/util/leaktest"
	"github.com/piotrszyma/sqlfmt/internal/util/log"
)

func TestContainsVars(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	testData := []struct {
		expr     string
		expected bool
	}{
		{`123`, false},
		{`123+123`, false},
		{`$1`, true},
		{`123+$1`, true},
		{`@1`, true},
		{`123+@1`, true},
	}

	for _, d := range testData {
		t.Run(d.expr, func(t *testing.T) {
			expr, err := parser.ParseExpr(d.expr)
			if err != nil {
				t.Fatalf("%s: %v", d.expr, err)
			}
			res := tree.ContainsVars(expr)
			if res != d.expected {
				t.Fatalf("%s: expected %v, got %v", d.expr, d.expected, res)
			}
		})
	}
}
