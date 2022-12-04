package access_to_fields_shorthands

import "fmt"

type A struct {
	FieldX int
}

func (a A) MethodA() {
	fmt.Print("Method A() called \n")
}

type B struct {
	*A
}

type C struct {
	B
}

func Run() {
	var c = &C{B: B{A: &A{FieldX: 5}}}

	// The following 4 lines are equivalent.
	_ = c.B.A.FieldX
	_ = c.B.FieldX
	_ = c.A.FieldX // A is a promoted field of C
	_ = c.FieldX   // FieldX is a promoted field

	// The following 4 lines are equivalent.
	c.B.A.MethodA()
	c.B.MethodA()
	c.A.MethodA()
	c.MethodA() // MethodA is a promoted method of C
}