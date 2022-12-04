package access_parent_fields_wrong_example

import "fmt"

type parent struct {
	x int
}

func (p parent) X() int {
	return p.x
}

type derivative struct {
	parent
	y int
}

func (d derivative) Y() int {
	return d.y
}

func (d derivative) SetX(x int) {
	d.parent.x = x
}

func newDerivative(x int, y int) derivative {
	return derivative{
		parent{x: x},
		y,
	}
}

func newDerivativePtr(x int, y int) *derivative {
	return &derivative{
		parent{x: x},
		y,
	}
}

func Run() {
	d := newDerivative(4, 5)
	fmt.Printf("X=%d Y=%d \n", d.x, d.y)
	d.SetX(10)
	// X will be the same (X=5)
	fmt.Printf("X=%d Y=%d \n", d.x, d.y)
}

func RunPtrExample() {
	d := newDerivativePtr(4, 5)
	fmt.Printf("X=%d Y=%d \n", d.x, d.y)
	d.SetX(10)
	// X will be the same (X=5)
	fmt.Printf("X=%d Y=%d \n", d.x, d.y)
}
