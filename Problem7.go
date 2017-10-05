// Niall Devery
// Adapted from: https://play.golang.org/p/5FUOzjBa-o

package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {// len gets the length of the string and splits it and for loop loops through the characters
		if input[i] != input[len(input)-i-1] {// if statement checks if the first and last letter are the same throughout the word if not true returns false
			return false
		}
	}
	return true // if the if statement doesn't find any exceptions the word is a palindrome
}

func main() {// 
	fmt.Println("Anna: ", isPalindrome("anna"))
	fmt.Println("Fluid: ", isPalindrome("fluid"))
	fmt.Println("Radar: ", isPalindrome("radar"))
	fmt.Println("Level: ", isPalindrome("level"))
	fmt.Println("Cheese: ", isPalindrome("cheese"))
	fmt.Println("Racecar: ", isPalindrome("racecar"))
}