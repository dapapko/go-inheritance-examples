package methodOverride

import "fmt"

type fooer struct{}

func (f fooer) Foo() {
	fmt.Println("Foo")
}

type notAnotherOneFooer struct {
	fooer
}

func (f notAnotherOneFooer) Foo() {
	fmt.Println("Not another one foo")
}

func Run() {
	f := fooer{}
	f.Foo()
	naof := notAnotherOneFooer{}
	naof.Foo()
}
