package basicInheritance

import "fmt"

type person struct {
	name string
}

func (p person) SetName(name string) {
	p.name = name
}

func (p person) Name() string {
	return p.name
}

func (p person) Print() {
	fmt.Printf("Employee name is %s \n", p.name)
}

type employee struct {
	person
}

func newEmployee(name string) employee {
	return employee{
		person{
			name: name,
		},
	}
}

func Run() {
	e := newEmployee("Sam")
	e.Print()
}
