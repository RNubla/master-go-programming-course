package main

import "fmt"

func main() {
	numberList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range numberList {
		if value%2 == 0 {
			fmt.Println(value, "Even")
		} else {
			fmt.Println(value, "Odd")
		}
	}
}
