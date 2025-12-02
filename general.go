package main

import (
	"fmt"
)

func CountAlpha(s string) int {
	count := 0

	for _, c := range s {
		// ---------------------------------------------------------
		// TODO:
		// Complete the condition below to check if `c` is an
		// alphabetical character.
		//
		// REQUIREMENTS:
		// - Count ONLY letters from A–Z and a–z
		// - Do NOT use unicode.IsLetter or any library functions
		// - Use ASCII comparisons only
		//
		// HINTS:
		// - Lowercase letters range from 'a' to 'z'
		// - Uppercase letters range from 'A' to 'Z'
		//
		// What you must fill in:
		//     if <your condition here> {
		//         count++
		//     }
		//
		// EXAMPLES:
		// 'A' → should be counted
		// 'g' → should be counted
		// '5' → NOT counted
		// '!' → NOT counted
		//
		// After completing the condition, the function should correctly
		// return how many letters appear in the string.
		// ---------------------------------------------------------

		// if ??? {
		//     count++
		// }
	}

	return count
}

func main() {
	// Example usage
	text := "Hello, World! 123"
	alphaCount := CountAlpha(text)
	fmt.Println("Number of alphabetic characters:", alphaCount)
}
