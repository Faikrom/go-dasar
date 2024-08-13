package main

import "fmt"

func main() {

	names := [...]string{"Fitroh", "ananda", "Ikrom", "Fargo", "Dani", "Lutfi"}

	slice1 := names[3:6]
	fmt.Println(slice1)

	slice2 := names[4:]
	fmt.Println(slice2)

	slice3 := names[:3]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

}
