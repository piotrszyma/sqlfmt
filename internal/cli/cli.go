// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package cli

import (
	"fmt"
	"strings"

	"github.com/piotrszyma/sqlfmt/internal/sql/sem/tree"
	// "github.com/piotrszyma/sqlfmt/internal/sql/parser"
)

// TODO(mjibson): This subcommand has more flags than I would prefer. My
// goal is to have it have just -e and nothing else. gofmt initially started
// with tabs/spaces and width specifiers but later regretted that decision
// and removed them. I would like to get to the point where we achieve SQL
// formatting nirvana and we make this an opinionated formatter with few or
// zero options, hopefully before 2.1.

// This file contains definitions for data types suitable for use by
// the flag+pflag packages.

// statementsValue is an implementation of pflag.Value that appends any
// argument to a slice.
type statementsValue []string

// Type implements the pflag.Value interface.
func (s *statementsValue) Type() string { return "statementsValue" }

// String implements the pflag.Value interface.
func (s *statementsValue) String() string {
	return strings.Join(*s, ";")
}

// Set implements the pflag.Value interface.
func (s *statementsValue) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// sqlfmtCtx captures the command-line parameters of the `sqlfmt` command.
// Defaults set by InitCLIDefaults() above.
type SqlfmtCtx struct {
	Len        int
	UseSpaces  bool
	TabWidth   int
	NoSimplify bool
	Align      bool
}

func RunSQLFmt(sqlfmtCtx SqlfmtCtx) error {
	if sqlfmtCtx.Len < 1 {
		return fmt.Errorf("line length must be > 0: %d", sqlfmtCtx.Len)
	}
	if sqlfmtCtx.TabWidth < 1 {
		return fmt.Errorf("tab width must be > 0: %d", sqlfmtCtx.TabWidth)
	}

	// var sl parser.Statements
	// in, err := ioutil.ReadAll(os.Stdin)
	// if err != nil {
	// 	return err
	// }
	// sl, err = parser.Parse(string(in))
	// if err != nil {
	// 	return err
	// }

	cfg := tree.DefaultPrettyCfg()
	cfg.UseTabs = !sqlfmtCtx.UseSpaces
	cfg.LineWidth = sqlfmtCtx.Len
	cfg.TabWidth = sqlfmtCtx.TabWidth
	cfg.Simplify = !sqlfmtCtx.NoSimplify
	cfg.Align = tree.PrettyNoAlign
	if sqlfmtCtx.Align {
		cfg.Align = tree.PrettyAlignAndDeindent
	}

	// for i := range sl {
	// 	fmt.Print(cfg.Pretty(sl[i].AST))
	// 	if len(sl) > 1 {
	// 		fmt.Print(";")
	// 	}
	// 	fmt.Println()
	// }

	return nil
}
