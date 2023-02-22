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
	x, y := 5, 10
	n := &x         // setting the pointer "n" to the address of "x"
	fmt.Println(*n) // while "n" is just pointing to the address of "x", when "*n" is called, it gives the exact value of the variable it is pointing to
	*n = 50
	fmt.Println(x)

	t := &y //assigning the memory address of "y" to "t", so "t" is pointer to "y"
	fmt.Println(*t)
	*t = 100       // this will change the value of the "y", because "*t" points to the value and not just the memory address
	fmt.Println(t) // this will only print the memory address
}
