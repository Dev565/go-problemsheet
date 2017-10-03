// Niall Devery


package main

import (
	"fmt"
	"time"
)


func main() {	
	// t used a the variable to hold time.now
	t := time.Now()
	// the hour and minute are then taken from time.now instead of all the information
	fmt.Println("The Time is "t.Hour(),":", t.Minute())
}