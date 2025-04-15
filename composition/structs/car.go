package structs

type engine struct {
	hp   int
	name string
}

func (e engine) HP() int {
	return e.hp
}

func (e engine) Name() string {
	return e.name
}

type wheels struct {
	wheelDimension int
	name           string
}

func (w wheels) WheelDimension() int {
	return w.wheelDimension
}

func (w wheels) Name() string {
	return w.name
}

// type Car struct {
// 	// this is called composition, have a struct inside a struct
// 		// this is a has-a relationship, not an is-a relationship
// 		Engine engine
// 		Wheels wheels
// 		CarName string
// }

// in the above case we need to put c.Engine.HP() to access the engine's HP method and c.Wheels.WheelDimension() to access the wheels' WheelDimension method
// but if we embed the struct inside the Car struct, we can access the fields directly like c.hp and c.wheelDimension

//######### we can declare this in another way ###########
type Car struct {
	CarName string
	engine
	wheels
	// this is called embedding, we can access the fields of the embedded struct directly
}

// func NewCar(carName string, hp int, wheelDimension int) Car {
// 	return Car{
// 		Engine: engine{hp: hp},
// 		Wheels: wheels{wheelDimension: wheelDimension},
// 		CarName: carName,
// 	}
// }

//## similarly this function can be written in another way ##
func NewCar(carName, wn, en string, hp, wheelDimension int) Car {
	return Car{
		CarName: carName,
		engine:  engine{hp: hp, name: en},
		wheels:  wheels{wheelDimension: wheelDimension, name: wn},
	}
}

//## now suppose I want to get the name of wheel and engine from the Car struct
//## so I will be using the same name for both of the structs, but we need to use the receiver to differentiate between the two functions
//## so we will be using the receiver to differentiate between the two functions

func (c Car) EngineName() string {
	return c.engine.Name()
}

func (c Car) WheelName() string {
	return c.wheels.Name()
}