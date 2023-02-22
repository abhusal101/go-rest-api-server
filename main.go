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
func (d *Dimension) Area() int {
	d.height = 8
	return d.length * d.width
}

func (d Dimension) Volume() int {
	d.height = 6
	return d.height * d.length * d.width
}

func main() {
	d := Dimension{10, 5, 6}
	fmt.Println(d.Area())
	fmt.Println(d) // after this line is run, we intend to get the output as {10 5 8} and we get it as desired, as we have used pointer to Dimension struct and that way we can manupilate the actual value
	fmt.Println(d.Volume())
	fmt.Println(d) // after this line is run, we intend to get the output as {10 5 6} but instead, it will still remain as {10 5 8}, because of Volume doesnt take the pointer to Dimesion
}
