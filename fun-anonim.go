package main

import "fmt"

type blackList func(string) bool

func registerUser(name string, blacklist blackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Fitroh", blacklist)
	registerUser("anjing", blacklist)
	registerUser("root", func(name string) bool {
		return name == "root"
	})

}
