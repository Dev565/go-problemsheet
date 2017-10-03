// Niall Devery
// Adapted from :https://play.golang.org/p/vhHmjhOMEo
package main

import "fmt"

func main() {
	// integer array to be sorted
	x := []int{
		90, 99, 37, 55,
		8, 42, 63, 81,
		26, 34, 12, 72,
		20, 45, 9, 17,
	}

	// 
	smallest, largest := x[0], x[0]
	for _, v := range x {
		if v > largest {
			largest = v
		}
		if v < smallest {
			smallest = v
		}
	}

	fmt.Println("The largest number is ", largest)
	fmt.Println("The smallest number is ", smallest)
}
