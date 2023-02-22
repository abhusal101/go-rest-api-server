package main

import "fmt"

type Dimension struct {
	length int
	width  int
	height int
}

// this is method in golang
// methods are the functions with recevier arguments

/*
(d Dimension) is the receiver type
Area() is the name of the method which does not take any arguments
and Area() return one int type
*/
func (d Dimension) Area() int {
	return d.length * d.width
}

func (d Dimension) Volume() int {
	return d.height * d.length * d.width
}

func main() {
	d := Dimension{10, 5, 6}
	fmt.Println("area: ", d.Area())
	fmt.Println("volume: ", d.Volume())
}
