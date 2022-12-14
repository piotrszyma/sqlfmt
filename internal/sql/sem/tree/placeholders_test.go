// Copyright 2019 The Cockroach Authors.
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
	"fmt"
	"testing"

	"github.com/piotrszyma/sqlfmt/internal/sql/types"
	"github.com/piotrszyma/sqlfmt/internal/util/leaktest"
	"github.com/piotrszyma/sqlfmt/internal/util/log"
)

func TestPlaceholderTypesIdentical(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	testCases := []struct {
		a, b      PlaceholderTypes
		identical bool
	}{
		{ // 0
			PlaceholderTypes{},
			PlaceholderTypes{},
			true,
		},
		{ // 1
			PlaceholderTypes{types.Int, types.Int},
			PlaceholderTypes{types.Int, types.Int},
			true,
		},
		{ // 2
			PlaceholderTypes{types.Int},
			PlaceholderTypes{types.Int, types.Int},
			false,
		},
		{ // 3
			PlaceholderTypes{types.Int, nil},
			PlaceholderTypes{types.Int, types.Int},
			false,
		},
		{ // 4
			PlaceholderTypes{types.Int, types.Int},
			PlaceholderTypes{types.Int, nil},
			false,
		},
		{ // 5
			PlaceholderTypes{types.Int, nil},
			PlaceholderTypes{types.Int, nil},
			true,
		},
		{ // 6
			PlaceholderTypes{types.Int},
			PlaceholderTypes{types.Int, nil},
			false,
		},
		{ // 7
			PlaceholderTypes{types.Int, nil},
			PlaceholderTypes{types.Int4, nil},
			false,
		},
		{ // 8
			PlaceholderTypes{types.Int},
			PlaceholderTypes{types.Int4},
			false,
		},
		{ // 9
			PlaceholderTypes{types.Int4},
			PlaceholderTypes{types.Int4},
			true,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := tc.a.Identical(tc.b)
			if res != tc.identical {
				t.Errorf("%v vs %v: expected %t, got %t", tc.a, tc.b, tc.identical, res)
			}
			res2 := tc.b.Identical(tc.a)
			if res != res2 {
				t.Errorf("%v vs %v: not commutative", tc.a, tc.b)
			}
		})
	}
}
