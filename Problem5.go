// Niall Devery
// https://www.rosettacode.org/wiki/Guess_the_number#Go

package main
 
import (
    "fmt"
    "math/rand"
    "time"
)
 
func main() {
    fmt.Print("Guess number from 1 to 10: ")
	// Random seed generated using the current time to determine the random number
    rand.Seed(time.Now().Unix())
    n := rand.Intn(10) + 1 // variable n used as a comparator to compare the users guess
    for guess := n; ; fmt.Print("No. Try again: ") { // for loop for users guesses
        switch _, err := fmt.Scan(&guess); { // switch statement to dtermine whether the guess is right 
        case err != nil:
            fmt.Println("\n", err, "\nSo, bye.") // catch for non-integer values
            return
        case guess == n:
            fmt.Println("Well guessed!") // if guess correctly programe ends
            return
        }
    }
}