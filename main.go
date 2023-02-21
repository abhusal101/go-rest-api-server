package main

import "fmt"

// function that takes string argument and also returns string value
func sayHello(name string) string {
	return "Hello, " + name
}

func main() {
	// calling the sayHello function along with passing the string value
	fmt.Println(sayHello("Asman"))
}
