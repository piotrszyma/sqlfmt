// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package pgerror_test

import (
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/piotrszyma/sqlfmt/internal/roachpb"
	"github.com/piotrszyma/sqlfmt/internal/sql/pgwire/pgcode"
	"github.com/piotrszyma/sqlfmt/internal/sql/pgwire/pgerror"
)

func TestWrap(t *testing.T) {
	testData := []struct {
		err error
	}{
		{errors.New("woo")},
		{&roachpb.TransactionRetryWithProtoRefreshError{}},
		{&roachpb.AmbiguousResultError{}},
	}

	for i, test := range testData {
		werr := pgerror.Wrap(test.err, pgcode.Syntax, "woo")

		if !errors.Is(werr, test.err) {
			t.Errorf("%d: original error not preserved; expected %+v, got %+v", i, test.err, werr)
		}
	}
}
