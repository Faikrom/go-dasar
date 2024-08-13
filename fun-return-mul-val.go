package main

import "fmt"

func getFullName() (string, string, string) {
	return "Fitroh", "ananda", "Ikrom"
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	fmt.Println("==============")

	name, _, _ := getFullName()
	fmt.Println(name)
}
