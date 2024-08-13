package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Tangerang"
	eko.Age = 25

	fmt.Println(eko)
	fmt.Println(eko.Name)

	fmt.Println("---------------")

	fitroh := Customer{
		Name:    "Fitroh",
		Address: "Tangerang",
		Age:     21,
	}
	fmt.Println(fitroh)

	fmt.Println("---------------")

	budi := Customer{"Budi", "Bandung", 25}
	fmt.Println(budi)

}
