# manyargs - 引数の多い関数を見つける

## 参考文献

- [14. 静的解析とコード生成 - Google スライド](https://docs.google.com/presentation/d/1I4pHnzV2dFOMbRcpA-XD0TaLcX6PBKpls6WxGHoMjOg/edit#slide=id.g80ffbfd5e3_0_168)
  - Greatな教科書。
- https://cs.opensource.google/go/x/tools/+/master:go/analysis/passes/
  - 本家リポジトリ。命名やディレクトリ構造などいろんことが参考になる。
- [GoAst Viewer](https://yuroyoro.github.io/goast-viewer/)
  - ASTノードの構造を見るのにたいへん役立つ。
  - 型パラメータの文法(Go1.18>)には対応していないので注意

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

