package polymorphism

type Shape interface {
	Render() string
}

type Circle struct {}

type Square struct {}

func (c Circle) Render() string {
	return "Rendering Circle"
}

func (s Square) Render() string {
	return "Rendering Square"
}