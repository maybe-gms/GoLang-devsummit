package main

import "fmt"

func main() {
	x := "Hello World"
	var y int = 230
	fmt.Println(x)
	fmt.Println("First number is", y)
	fmt.Printf("%v  Type:%T\n", y, y) // Using %T to print the type of y
}
