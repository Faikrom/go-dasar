package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke ", i)
	}

	fmt.Println("===================")

	for i := 0; i < 10; i++ {
		fmt.Println("perulangan ke ", i)
		if i == 5 {
			break
		}
	}
}
