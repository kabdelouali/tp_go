//Create more sophisticated but still simple programs using Go and data structure usage.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")

	fmt.Println("ParsePhone() test")
	fmt.Printf("ParsePhone(%q) = %q\n", "123-456-7890", ParsePhone("123-456-7890"))
	fmt.Printf("ParsePhone(%q) = %q\n", "1 2 3 4 5 6 7 8 9 0", ParsePhone("1 2 3 4 5 6 7 8 9 0"))

	fmt.Println("Anagram() test")
	fmt.Printf("Anagram(%q, %q) = %v\n", "12345", "52314", Anagram("12345", "52314"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "21435", "53241", Anagram("21435", "53241"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "12346", "52314", Anagram("12346", "52314"))
	fmt.Printf("Anagram(%q, %q) = %v\n", "123456", "52314", Anagram("123456", "52314"))

	fmt.Println("FindEvens() test")
	fmt.Printf("FindEvens(%v) = %v\n", []int{1, 2, 3, 4}, FindEvens([]int{1, 2, 3, 4}))

	fmt.Println("SliceProduct() test")
	fmt.Printf("SliceProduct(%v) = %v\n", []int{5, 6, 8}, SliceProduct([]int{5, 6, 8}))

	fmt.Println("Unique() test")
	fmt.Printf("Unique(%v) = %v\n", []int{1, 2, 3, 4, 4, 5, 6, 6}, Unique([]int{1, 2, 3, 4, 4, 5, 6, 6}))

	fmt.Println("InvertMap() test")
	fmt.Printf("InvertMap(%v) = %v\n", map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}, InvertMap(map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}))
}

// ParsePhone parses a string of numbers into the format 06 22 14 33 44.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "12 34 56 78 90"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "12 34 56 78 90"
func ParsePhone(phone string) string {
	res := ""
	cpt := 0
	for i := 0; i < len(phone); i++ {
		_, err := strconv.Atoi(string(phone[i]))
		if err == nil {
			cpt++
			res += string(phone[i])
		}
		if cpt == 2 {
			if len(res) < 14 {
				res += " "
				cpt = 0
			}
		}
	}
	return res
}

// Write a function to check whether two given strings are anagram of each other or not.
// An anagram of a string is another string that contains same characters,
// only the order of characters can be different. For example, “abcd” and “dabc” are anagram of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	} else {
		for i := 0; i < len(s1); i++ {
			if strings.Contains(s2, string(s1[i])) == false {
				return false
			}
		}
	}
	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var res []int
	for i := 0; i < len(e); i++ {
		if e[i]%2 == 0 {
			res = append(res, e[i])
		}
	}
	return res
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	res := 1
	for i := 0; i < len(e); i++ {
		res = res * e[i]
	}
	return res
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	var res []int
	for i := 0; i < len(e); i++ {
		add := true
		for j := 0; j < len(e); j++ {
			if e[i] == e[j] && i != j {
				add = false
			}
		}
		if add == true {
			res = append(res, e[i])
		}
	}
	return res
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	res := make(map[int]string)
	for key, element := range kv {
		res[element] = key
	}
	return res
}
