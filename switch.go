package main

import (
	"fmt"
)

func main() {

	nama := "Fitrohananada Ikrom"

	switch nama {
	case "Fitroh":
		fmt.Println("Hallo, Fitroh")
	case "Dani":
		fmt.Println("Hallo, Dani")
	default:
		fmt.Println("Hallo, Strangers")
	}

	fmt.Println("========================")

	// switch length := len(nama); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	fmt.Println("========================")

	length := len(nama)

	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
