package main

import "fmt"

func dimensions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = length * width * height
	return
}

func main() {
	area, volume := dimensions(5, 5, 5)
	fmt.Println("area: ", area, ", volume:", volume)
}
