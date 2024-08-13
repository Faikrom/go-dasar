package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Fitroh"
	name[1] = "ananda"
	name[2] = "ikrom"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var nama = [3]int{
		80,
		90,
		100,
	}

	fmt.Println(nama)

	fmt.Println("==========================")

	var nama2 = [...]int{
		80,
		90,
		100,
		90,
		50,
	}

	fmt.Println(len(nama2))
	fmt.Println(nama2)

	

}
