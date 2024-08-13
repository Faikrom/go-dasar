package main

import "fmt"

func main() {
	var name = "Fitroh"

	if name == "Fitroh" {
		fmt.Println("Hello Fitroh")
	} else if name == "Dani" {
		fmt.Println("Hello Dani")
	} else {
		fmt.Println("Hello Stranger")
	}

	fmt.Println("======================")

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	fmt.Println("======================")

}
