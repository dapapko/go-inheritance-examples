package main

import (
	"fmt"
	access_parent_fields "go-examples/access-parent-fields"
	access_parent_fields_wrong_example "go-examples/access-parent-fields-wrong-example"
	basicInheritance "go-examples/basic-inheritance"
	methodOverride "go-examples/method-override"
	multipleInheritance "go-examples/multiple-inheritance"
)

func main() {
	fmt.Println("===Basic inheritance example===")
	basicInheritance.Run()
	fmt.Println("===Override parent method===")
	methodOverride.Run()
	fmt.Println("===Multiple inheritance example===")
	multipleInheritance.Run()
	fmt.Println("===Parent fields initialization and access===")
	access_parent_fields.Run()
	fmt.Println("===Parent fields initialization and access (ptr example)===")
	access_parent_fields.RunPtrExample()
	fmt.Println("===Parent fields initialization and access (wrong example)===")
	access_parent_fields_wrong_example.Run()
	fmt.Println("===Parent fields initialization and access (wrong ptr example)===")
	access_parent_fields_wrong_example.RunPtrExample()
	fmt.Println("===Access to parents fields and methodas shorthands example===")
	access_to_fields_shorthands.RunPtrExample()
}
