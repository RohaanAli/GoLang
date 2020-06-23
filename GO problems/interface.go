package main

import "fmt"


type dimension interface{
	Area() float64
	peri() float64
}

func main() {
	shape1 := shape{5.0,5.0}
	fmt.Println(shape1.Area())
	fmt.Println(shape1.peri())
}

type shape struct {
	length float64 
	width float64  
}

func (s shape) Area()float64{
	len := s.length
	wid := s.width

	return len * wid
}

func (sh shape) peri() float64{
	return 2.0 * (sh.length  + sh.width)
}



