package main

import "os"

var exit bool

func atoip(s string) int {
	var n int
	for _, digit := range s {
		if digit > '9' || digit < '0' {
			exit = true
		}
	}
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	if n == 0 {
		exit = true
	}
	return n
}

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}
	n1 := atoip(os.Args[1])
	n2 := atoip(os.Args[2])
	var gcd int
	if exit {
		return
	}

	for i := 1; i <= n2; i++ {
		if n1%i == 0 && n2%i == 0 {
			gcd = i
		}
	}

	// -----------------------------------------------------
	// TODO: Convert gcd (which is an integer)
	//       into a string WITHOUT using:
	//          - strconv
	//          - fmt
	//          - Sprintf
	//
	// HOW TO DO IT:
	// 1. Create an empty string variable (e.g., result := "").
	// 2. While gcd is greater than 0:
	//        - Extract the last digit using: digit := gcd % 10
	//        - Convert that digit to a character by adding '0'
	//              char := rune(digit + '0')
	//        - Prepend the character to the result string
	//              result = string(char) + result
	//        - Remove the last digit from gcd:
	//              gcd = gcd / 10
	//
	// 3. After the loop ends, write the string to stdout:
	//        os.Stdout.WriteString(result + "\n")
	//
	// IMPORTANT RULE:
	// The digits must be added to the FRONT of the string,
	// not the back, because you extract them from right to left.
	//
	// EXAMPLE:
	// gcd = 12
	// step 1 -> digit = 2  → result = "2"
	// step 2 -> digit = 1  → result = "12"
	//
	// If you append instead of prepend, you would get "21".
	// -----------------------------------------------------
}
