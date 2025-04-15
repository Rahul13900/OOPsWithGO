package main

import (
	"fmt"
	polymorphism "oops/polymorphism/shape"
)

func main() {
	var c polymorphism.Shape = polymorphism.Circle{}
	var s polymorphism.Shape = polymorphism.Square{}
	fmt.Println(c.Render(), s.Render())
}
