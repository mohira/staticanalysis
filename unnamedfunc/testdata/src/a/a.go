package a

func f(int, int) {} // want "f 名前無し引数を使った関数"

type _ interface {
	f(int, int) // interfaceは対象外
}

type T int

func (t *T) f(int, int) {} // メソッドは対象外
