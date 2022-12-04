package multipleInheritance

import "fmt"

type fooer struct{}

func (f fooer) Foo() {
	fmt.Println("Foo")
}

type barer struct{}

func (b barer) Bar() {
	fmt.Println("Bar")
}

type fooerBarer struct {
	fooer
	barer
}

func Run() {
	fb := fooerBarer{}
	fb.Bar()
	fb.Foo()
}
