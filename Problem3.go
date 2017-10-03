// Niall Devery

package main

import "fmt"

func main() {
	
	var count int = 0

	for count<100{
		count++
		//logical & to check if id divisable by both 3 and 5
	   if count % 3 == 0 && count % 5 ==0 { 
		   fmt.Println("FizzBuzz")
       } else if count % 3 == 0 {
		   fmt.Println("Fizz")
	   } else if count % 5 == 0 {
		   fmt.Println("Buzz")
	   } else {
		   fmt.Println(count)
	   }//end if else

	}//end for count<=100

}//end main