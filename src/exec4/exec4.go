//Use interfaces and data structures to build programs
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("========== Problem 1: Sorting Names ==========")
	people := PersonSlice{NewPerson("Raphel", "Chab"), NewPerson("Julien", "Chab"), NewPerson("Zhe", "Liu"), NewPerson("zhe", "Liu")}
	fmt.Printf("Before: %v\n", people)
	sort.Sort(PersonSlice(people))
	fmt.Printf("After: %v\n", people)

	fmt.Println("========== Problem 2: Functional Programming ==========")
	sum := func(x, y int) int { return x + y }
	fmt.Printf("Fold([]int{1, 2, 3, 4, 5}, 0, add) = %v\n", Fold([]int{1, 2, 3, 4, 5}, 0, sum))
	multiplication := func(x, y int) int { return x * y }
	fmt.Printf("Fold([]int{1, 2, 3, 4, 5}, 1, mult) = %v\n", Fold([]int{1, 2, 3, 4, 5}, 1, multiplication))
}

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type PersonSlice []*Person

var last_ID = 0

func (ps PersonSlice) Len() int {
	return len([]*Person(ps))
}

func (ps PersonSlice) Less(i, j int) bool {
	if c := strings.Compare(ps[i].LastName, ps[j].LastName); c != 0 {
		return c < 0
	}
	if c := strings.Compare(ps[i].FirstName, ps[j].FirstName); c != 0 {
		return c < 0
	}
	return ps[i].ID < ps[j].ID
}

func (a PersonSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (p *Person) String() string {
	return fmt.Sprintf("%s %s(%v)", p.FirstName, p.LastName, p.ID)
}

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
func NewPerson(first, last string) *Person {
	person := new(Person)
	person.FirstName = first
	person.LastName = last
	person.ID = last_ID + 1
	last_ID++
	return person
}

// Problem 2: Functional Programming
// Write a function Fold which applies a function repeatedly on a slice,
// producing a single value via repeated application of an input function.
// The behavior of Fold should be as follows:
//   - When s is empty, return v (default value)
//   - When s has 1 value (x0), apply f once: f(v, x0)
//   - When s has 2 values (x0, x1), apply f twice, from left to right: f(f(v, x0), x1)
//   - Continue this pattern recursively to obtain the final result.

// Fold applies a left to right application of f on s starting with v.
// Note the argument signature of f - func(int, int) int.
// This means f is a function which has 2 int arguments and returns an int.
func Fold(s []int, v int, f func(int, int) int) int {
	if len(s) == 0 {
		return v
	} else {
		v_prime := f(v, s[0])
		return Fold(s[1:], v_prime, f)
	}
}
