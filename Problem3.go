// Niall Devery

package main

import "fmt"

func main() {
	
	var count int = 0

	for count<100{
		count++ // count incrementation each time the loops is passed
	   if count % 3 == 0 && count % 5 ==0 {// logical & to check if id divisable by both 3 and 5 (must be first condition checked
		   fmt.Println("FizzBuzz")
       } else if count % 3 == 0 {// condition if value is divisible by 3 to print fizz
		   fmt.Println("Fizz")
	   } else if count % 5 == 0 {// condition if value is divisable by 5 to print buzz
		   fmt.Println("Buzz")
	   } else {// prints the number if no conditions are met
		   fmt.Println(count)
	   }//end if else

	}//end for count<=100

}//end main