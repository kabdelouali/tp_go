// package main

// import "fmt"

// func main() {
// 	fmt.Printf("hello, world\n")
// }

//become familiar writing Go programs and running them on your machine

package main

import (
	"fmt"
)

func main() {
	// use the main function for testing your functions
	fmt.Println("Hello, world!")

	fmt.Println("Fizzbuzz() test")
	fmt.Printf("Fizzbuzz(%v) = %v\n", 27, Fizzbuzz(27))
	fmt.Printf("Fizzbuzz(%v) = %v\n", 25, Fizzbuzz(25))
	fmt.Printf("Fizzbuzz(%v) = %v\n", 105, Fizzbuzz(105))

	fmt.Println("IsPrime() test")
	fmt.Printf("IsPrime(%v) = %v\n", 1, IsPrime(1))
	fmt.Printf("IsPrime(%v) = %v\n", 2, IsPrime(2))
	fmt.Printf("IsPrime(%v) = %v\n", 22, IsPrime(22))
	fmt.Printf("IsPrime(%v) = %v\n", 32, IsPrime(32))

	fmt.Println("IsPalindrome() test")
	fmt.Printf("IsPalindrome(%v) = %v\n", "abccba", IsPalindrome("abccba"))
	fmt.Printf("IsPalindrome(%v) = %v\n", "abccbb", IsPalindrome("abccbb"))
	fmt.Printf("IsPalindrome(%v) = %v\n", "abcbba", IsPalindrome("abcbba"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	} else {
		return ""
	}
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
