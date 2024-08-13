package main

import "fmt"

func main() {

	person := map[string]string{
		"nama":    "Fitroh",
		"address": "Tangerang",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["nama"])

	var book = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Semangat"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
