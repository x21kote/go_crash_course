package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p *Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Samansa", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Thompson")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
