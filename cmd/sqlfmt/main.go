package main

import (
	"fmt"
	"log"

	"github.com/piotrszyma/sqlfmt/internal/cli"
)

func main() {
	err := cli.RunSQLFmt(cli.SqlfmtCtx{
		Len:        120,
		UseSpaces:  true,
		TabWidth:   2,
		NoSimplify: true,
		Align:      true,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("formatting failed: %w", err))
	}

	log.Print("Done")
}
