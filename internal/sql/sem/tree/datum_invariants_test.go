// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tree

import (
	"testing"

	"github.com/piotrszyma/sqlfmt/internal/sql/types"
	"github.com/piotrszyma/sqlfmt/internal/util/leaktest"
	"github.com/piotrszyma/sqlfmt/internal/util/log"
)

func TestAllTypesCastableToString(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	for _, typ := range types.Scalar {
		if err := resolveCast("", typ, types.String, true); err != nil {
			t.Errorf("%s is not castable to STRING, all types should be", typ)
		}
	}
}

func TestAllTypesCastableFromString(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	for _, typ := range types.Scalar {
		if err := resolveCast("", types.String, typ, true); err != nil {
			t.Errorf("%s is not castable from STRING, all types should be", typ)
		}
	}
}
