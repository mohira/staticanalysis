# manyargs - 引数の多い関数を見つける

## モチベーション: なぜ引数の多い関数を見つけたいのか？

設計改善のヒントになるから。引数が多い関数は設計が甘い可能性が高い。

## 仕様

[testdata/src/a/a.go](./testdata/src/a/a.go)を見てください。

## `go vet`で呼び出す

```
$ pwd
staticanalysis/manyargs/cmd/manyargs

$ go build .

$ go vet -vettool=manyargs fmt
# fmt
/usr/local/go/src/fmt/format.go:194:1: 引数が多いな！
```

