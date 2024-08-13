package main

import "fmt"

func main() {
	name := "Fitroh"
	counter := 0

	increment := func() {
		// name = "Budi" salah karna akan mengganti name diatas
		name := "Budi"
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
