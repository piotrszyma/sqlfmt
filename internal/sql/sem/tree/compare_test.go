// Copyright 2017 The Cockroach Authors.
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
	"context"
	"testing"

	"github.com/piotrszyma/sqlfmt/internal/settings/cluster"
	"github.com/piotrszyma/sqlfmt/internal/sql/sem/eval"
	"github.com/piotrszyma/sqlfmt/internal/sql/sem/tree"
	"github.com/piotrszyma/sqlfmt/internal/sql/sem/tree/treecmp"
	"github.com/piotrszyma/sqlfmt/internal/sql/types"
	"github.com/piotrszyma/sqlfmt/internal/util/leaktest"
	"github.com/piotrszyma/sqlfmt/internal/util/log"
)

func TestEvalComparisonExprCaching(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	testExprs := []struct {
		op          treecmp.ComparisonOperatorSymbol
		left, right string
		cacheCount  int
	}{
		// Comparisons.
		{treecmp.EQ, `0`, `1`, 0},
		{treecmp.LT, `0`, `1`, 0},
		// LIKE and NOT LIKE
		{treecmp.Like, `TEST`, `T%T`, 1},
		{treecmp.NotLike, `TEST`, `%E%T`, 1},
		// ILIKE and NOT ILIKE
		{treecmp.ILike, `TEST`, `T%T`, 1},
		{treecmp.NotILike, `TEST`, `%E%T`, 1},
		// SIMILAR TO and NOT SIMILAR TO
		{treecmp.SimilarTo, `abc`, `(b|c)%`, 1},
		{treecmp.NotSimilarTo, `abc`, `%(b|d)%`, 1},
		// ~, !~, ~*, and !~*
		{treecmp.RegMatch, `abc`, `(b|c).`, 1},
		{treecmp.NotRegMatch, `abc`, `(b|c).`, 1},
		{treecmp.RegIMatch, `abc`, `(b|c).`, 1},
		{treecmp.NotRegIMatch, `abc`, `(b|c).`, 1},
	}
	for _, d := range testExprs {
		expr := &tree.ComparisonExpr{
			Operator: treecmp.MakeComparisonOperator(d.op),
			Left:     tree.NewDString(d.left),
			Right:    tree.NewDString(d.right),
		}
		ctx := eval.NewTestingEvalContext(cluster.MakeTestingClusterSettings())
		defer ctx.Mon.Stop(context.Background())
		ctx.ReCache = tree.NewRegexpCache(8)
		typedExpr, err := tree.TypeCheck(context.Background(), expr, nil, types.Any)
		if err != nil {
			t.Fatalf("%v: %v", d, err)
		}
		if _, err := eval.Expr(ctx, typedExpr); err != nil {
			t.Fatalf("%v: %v", d, err)
		}
		if typedExpr.(*tree.ComparisonExpr).Op.EvalOp == nil {
			t.Errorf("%s: expected the comparison function to be looked up and memoized, but it wasn't", expr)
		}
		if count := ctx.ReCache.Len(); count != d.cacheCount {
			t.Errorf("%s: expected regular expression cache to contain %d compiled patterns, but found %d", expr, d.cacheCount, count)
		}
	}
}
