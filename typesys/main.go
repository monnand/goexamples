package main

import (
	"fmt"
)

type Duck struct {
	Name string
}

func (d *Duck) Eat() {
	fmt.Println(d.Name, "Duck is having dinner!")
}

type DonaldDuck struct {
	Duck
	Age int
}

type Animal interface {
	Eat()
	Drink(amount int) bool
}

func (d *Duck) Drink(amount int) bool {
	if amount < 100 {
		fmt.Println("That's not enought!")
		return false
	}
	fmt.Println("OK, I'm full")
	return true
}

func main() {
	var a Animal
	d := new(Duck)
	d.Name = "Don"
	a = d
	a.Eat()
}

