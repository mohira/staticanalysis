package a

// 捕捉対象: 関数宣言
// ast.Decl > ast.FuncDecl > ast.FuncType > ast.FieldList
func f1(a int, b int, c int, d int, e int) {} // want "引数が多いな！"
func f2(a, b, c, d, e int)                 {} // want "引数が多いな！"
func f3(a, b int, c, d int, e int)         {} // want "引数が多いな！"
func f4(int, int, int, int, int)           {} // want "引数が多いな！"

// 捕捉対象: 関数リテラル
// ast.GenDecl > ast.Spec > ast.ValueSpec > ast.Expr > ast.FuncLit > ast.FuncType > ast.FieldList
var _ = func(a int, b int, c int, d int, e int) {} // want "引数が多いな！"

// 捕捉対象: 構造体のフィールド
// ast.GenDecl > ast.Spec > ast.TypeSpec > ast.StructType > ast.Filed > ast.FuncType > ast.FieldList
type _ struct {
	manyArgsX func(a int, b int, c int, d int, e int) // want "引数が多いな！"
	manyArgsY func(a, b int, c, d int, e int)         // want "引数が多いな！"
}

// 捕捉対象: インタフェース
// ast.GenDecl > ast.Spec > ast.TypeSpec > ast.InterfaceType > ast.Field > ast.FuncType > ast.FieldList
type _ interface {
	manyArgsA(a int, b int, c int, d int, e int) // want "引数が多いな！"
	manyArgsB(a, b int, c, d int, e int)         // want "引数が多いな！"
}

// 捕捉対象ではない: 可変長引数は1つとカウントする
func _(a, b, c int, d ...int) {}

// 捕捉対象ではない: 型パラメータ
// ast.FuncDecl > ast.FuncType > TypeParams(*ast.FieldList)
func _[a int, b int, c int, d int, e int](x, y, z int) {}
