package main

import (
	"fmt"
	"oops/encapsulation/structs"
)

func main() {
	// p := structs.Person{}
	// another way to create a Person Object
	p := structs.NewPerson("John", "Doe", 30)

	// p.firstName = "John" // cannot set firstName because it is private
	// therefore we need to use the receiver function to set the first name
	p.SetFirstName("John")
	fmt.Println(p)
}
