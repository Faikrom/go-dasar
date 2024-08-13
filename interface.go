package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	fitroh := Person{
		Name: "Fitroh",
	}

	SayHello(fitroh)

	cat := Animal{
		Name: "Cat",
	}

	SayHello(cat)

	var dog Animal
	dog.Name = "Dog"
	SayHello(dog)
}
