package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke : ", counter)
		counter++
	}

	fmt.Println("=================")

	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke : ", i)
	}

	fmt.Println("=================")

	slice := []string{"Fitroh", "ananda", "Ikrom"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("=================")

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	fmt.Println("=================")

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Println("=================")

	person := make(map[string]string)
	person["name"] = "Fitroh"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
