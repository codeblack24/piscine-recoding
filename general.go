package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	for player := 0; player <= 4; player++ {
		fmt.Printf("Player: ", player)

		start := player * 4
		end := start + 2

		for i := start; i <= end; i++ {
			if i == end+1 {
				fmt.Printf("%d", deck[i])
			} else {
				fmt.Printf("%d, ", deck[start])
			}
		}

		z01.PrintRune('\n')
		z01.PrintRune('\n')
	}
}
