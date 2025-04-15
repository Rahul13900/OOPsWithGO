package main

import (
	"fmt"
	"oops/composition/structs"
)

func main() {
	c := structs.NewCar("Toyota","C2", "MHawk", 150, 16)
	fmt.Println(c.EngineName())
}
