package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	arrInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, intValue := range arrInt {
		if intValue%2 == 0 {
			fmt.Println(intValue, " is even")
		} else {
			fmt.Println(intValue, " is odd")
		}
	}
}
