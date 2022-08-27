# sqlfmt

Goal of this project is to extract `sqlfmt` command from `cockroach` and
provide CLI interface similar to formatters like `black` or `golangci-lint`.

TODO:
- [ ] Figure out a way to run only `sqlfmt`
- [ ] (maybe) extract parser from `cockroach`
- [ ] (maybe) extract pretty printer from `cockroach`
- [ ] (maybe) make it buildable using `go build` system instead of `BLAZE`


## Reformat SQL

To format `*.sql` files in given directory `/dir` in place, run:
```
sqlfmt --write /dir
```

## Check SQL formatting (for CI)

To only check formatting and return non-zero status if some file needs reformatting, run:
```
sqlfmt --check /dir
```

## This tool requires Go 1.17

To install:
1. https://go.dev/dl/go1.17.13.linux-amd64.tar.gz
