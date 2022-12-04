package ambiguous_call

import "fmt"

type fooer struct {
	X int
}

func (f fooer) Foo() {
	fmt.Println("Foo")
}

type anotherFooer struct {
	Y int
}

func (a anotherFooer) Foo() {
	fmt.Println("Bar")
}

type ambiguousFooer struct {
	fooer
	anotherFooer
	Z int
}

func Run() {
	f := ambiguousFooer{}
	// compilation error
	f.Foo()
}
