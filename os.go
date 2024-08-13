package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Argument : ")
	fmt.Println(args)

	fmt.Println(args[1])
	fmt.Println(args[2])

	fmt.Println("---------------")

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error : ", err.Error())
	}

	fmt.Println("---------------")

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
