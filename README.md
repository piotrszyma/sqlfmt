# sqlfmt

TODO:
- [ ] extract parser
- [ ] extract pretty printer

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
