// Begin gaining familiarity with Go-style concurrency primitives.
// You will need to refer to the Go documentation.
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("========== Problem 1 : File processing ==========")
	Sum("numbers.txt", "sum.txt")

	fmt.Println("\n========== Problem 2 :  Concurrent map access==========")
	d := EnseirbDirectory{
		directory: make(map[int]string),
	}
	total := 1001
	var wg sync.WaitGroup
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			switch i % 3 {
			case 0:
				d.Add(1, "Aurore LI")
			case 1:
				d.Add(2, "Catie")
			case 2:
				d.Remove(2)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("d.directory = %v\n", d.directory)
}

// Problem 1: File processing with interfaces
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// You should use the interfaces for io.Reader and
// io.Writer to do this.

// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.
// You should expect your input to end with a newline, and the output should
// have a newline after the result
func Sum(input, output string) {
	file, err := os.Open(input)
	if err != nil {
		fmt.Print(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Print(err)
	}
	sum := 0
	var str string
	for i := 0; i < count; i++ {
		if string(data[i]) != "\n" {
			str += string(data[i])
		} else {
			r, _ := strconv.Atoi(str)
			sum += r
			str = ""
		}
	}
	fmt.Print(sum)
	output_file, _ := os.Create(output)
	output_file.Write([]byte(strconv.Itoa(sum)))
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// You will build a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// EnseirbDirectory is a mapping from ID number to name (12345678 -> "Aurore LI").
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex
type EnseirbDirectory struct {
	mutex     sync.RWMutex
	directory map[int]string
}

// Add inserts a new student to the Enseirb Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *EnseirbDirectory) Add(id int, name string) {
	d.mutex.Lock()
	d.directory[id] = name
	d.mutex.Unlock()
}

// Get fetches a student from the Enseirb Directory by their ID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *EnseirbDirectory) Get(id int) string {
	d.mutex.RLock()
	value := d.directory[id]
	d.mutex.RUnlock()
	return value
}

// Remove a student to the Enseirb Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *EnseirbDirectory) Remove(id int) {
	d.mutex.Lock()
	delete(d.directory, id)
	d.mutex.Unlock()
}
