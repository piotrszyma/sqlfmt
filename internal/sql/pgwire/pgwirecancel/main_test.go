// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package pgwirecancel_test

import (
	"os"
	"testing"

	"github.com/piotrszyma/sqlfmt/internal/security/securityassets"
	"github.com/piotrszyma/sqlfmt/internal/security/securitytest"
	"github.com/piotrszyma/sqlfmt/internal/server"
	"github.com/piotrszyma/sqlfmt/internal/testutils/serverutils"
	"github.com/piotrszyma/sqlfmt/internal/testutils/testcluster"
	"github.com/piotrszyma/sqlfmt/internal/util/randutil"
)

//go:generate ../../../util/leaktest/add-leaktest.sh *_test.go

func TestMain(m *testing.M) {
	securityassets.SetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)
	os.Exit(m.Run())
}
