package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(90, 80, 50, 70, 10, 30)
	fmt.Println(total)

	slice := []int{10, 60, 70, 40}
	total = sumAll(slice...)
	fmt.Println(total)
}
