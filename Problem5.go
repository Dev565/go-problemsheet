// Niall Devery
// http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html

package main
import(
    "fmt"
    "math/rand"
    "time"
)

//this generates random number between given range
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := xrand(1, 10)
    guessTaken := 0
    var guess int

    
    //this is the while loop
    for guessTaken < 10 {
        fmt.Println("Take a guess...")
        fmt.Scanf("%d", &guess)
		fmt.Scanf("%d")// extra scanf to prevent the prompt printing twice
        guessTaken++// incremtation of guesses taken
		// if statements to determine whether the guess is higher or lower than the random number
		// catch for characters 
//		if guess != nil{
//			fmt.Println("None interger value entered")
//		}
        if guess < myrand {
            fmt.Println("Your guess is too low.")
        }
        if guess > myrand {
            fmt.Println("Your guess is too high.")
        }
        if guess == myrand {// if the number is guessed the loop broken
            break
        }
    }
    if guess == myrand {// if the name was correct the user is notified and the programme exits
        fmt.Printf("You guessed the number in %d tries\n", guessTaken)
    } else {
        fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
    }
}