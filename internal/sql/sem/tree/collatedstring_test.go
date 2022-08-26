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
	"github.com/piotrszyma/sqlfmt/internal/sql/types"
	"github.com/piotrszyma/sqlfmt/internal/util/leaktest"
	"github.com/piotrszyma/sqlfmt/internal/util/log"
)

func TestCastToCollatedString(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	cases := []struct {
		typ      *types.T
		contents string
	}{
		{types.MakeCollatedString(types.String, "de"), "test"},
		{types.MakeCollatedString(types.String, "en"), "test"},
		{types.MakeCollatedString(types.MakeString(5), "en"), "test"},
		{types.MakeCollatedString(types.MakeString(4), "en"), "test"},
		{types.MakeCollatedString(types.MakeString(3), "en"), "tes"},
	}
	ctx := context.Background()
	for _, cas := range cases {
		t.Run("", func(t *testing.T) {
			expr := &tree.CastExpr{
				Expr:       tree.NewDString("test"),
				Type:       cas.typ,
				SyntaxMode: tree.CastShort,
			}
			semaCtx := tree.MakeSemaContext()
			typedexpr, err := expr.TypeCheck(ctx, &semaCtx, types.Any)
			if err != nil {
				t.Fatal(err)
			}
			evalCtx := eval.NewTestingEvalContext(cluster.MakeTestingClusterSettings())
			defer evalCtx.Stop(context.Background())
			val, err := eval.Expr(evalCtx, typedexpr)
			if err != nil {
				t.Fatal(err)
			}
			switch v := val.(type) {
			case *tree.DCollatedString:
				if v.Locale != cas.typ.Locale() {
					t.Errorf("expected locale %q but got %q", cas.typ.Locale(), v.Locale)
				}
				if v.Contents != cas.contents {
					t.Errorf("expected contents %q but got %q", cas.contents, v.Contents)
				}
			default:
				t.Errorf("expected type *DCollatedString but got %T", v)
			}
		})
	}
}
