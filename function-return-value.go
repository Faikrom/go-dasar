package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello" + name
	}
}

func main() {
	result := getHello("Fitroh")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
