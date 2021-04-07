package main

import "fmt"

func main() {
	// using var
	// var name string = "Brad"
	var age int32 = 37
	const isCool = true
	var size float32 = 2.3

	// Shorthand
	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
