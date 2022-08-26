# sqlfmt

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

