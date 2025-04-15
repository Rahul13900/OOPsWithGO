package structs

// to make the struct private, start with a lowercase letter
// to make the struct public, start with an uppercase letter
type Person struct {
	firstName string
	lastName  string
	age       int
}

// its kind of constructor function
// there is no constructor in goLang, but we can create a function that returns a struct
func NewPerson(firstName, lastName string, age int) Person {
	// we can also write validation logic here
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

// functions specific to the struct are called receiver functions
func (p Person) Walk() string {
	return p.firstName + " " + p.lastName + " likes walking"
}

// pointer based receiver function
func (p *Person) SetFirstName(firstName string) {
	p.firstName = firstName
}
