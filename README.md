# go-problemsheet
MY solution to go problems

Problem 1

Write a program that prints Hello, world! in Japanese (using Japanese characters) to the screen.

For this problem I just used a fmt.Println along with hello world in english and in japanese.

Problem 2

Write a program that prints the current time and date to the console.

For this problem I used the import time to get the time.Now and from there using .hour and .minute to use only 
the information I wanted to print to the screen.

Problem 3

Write a program that prints the numbers from 1 to 100, each on a new line, to the console, except for the 
following conditions. For multiples of three print Fizz instead of the number, and for multiples of five print
Buzz. For numbers that are multiples of both three and five print FizzBuzz.

Using a for loop, if, and else if statements I incremented the count variable to 100. Then using a logical 
& to check if the number was divisable by both 3 and 5 for to be sure that I didnt just print "Fizz" when I 
should be printing "FizzBuzz". If the number was not divisable by 3 or 5 the number would just be printed.

Problem 5

Write a guessing game where the user must guess a randomly generated number. After every guess the program 
tells the user whether their number was too high or too low. At the end, the number of tries should be printed.
It counts only as one try if they input the same number multiple times consecutively.

I found useful resources online and decided to use them as a starting point after understanding the code.
The programme gets a random number from between 1 and 10 using the current time from the "time" and "math/rand"
imports. When i first ran the code it had printed the prompt to the user twice and had taken a guess to counter
that I added a second scanf @ line 27. Using if statements I tried to catch if the user had entered a character
instead of a number using "nil" but couldn't get it to work.

Problem 6

Write a function that returns the largest and smallest elements in a list.

Using an un-sorted integer array along with a for loop and if statements I cycled through the array determining
whether the integer value was bigger or smaller the the range

Problem 7

Write a function that tests whether a string is a palindrome. A palindrome is a string that reads the same 
in reverse, e.g. radar.

I found resources online that were very similar to what I wanted to do. Using a boolean operator to dtermine
whether a word was a palindrome or not the programme gets the length of the string and determines whether the 
last and the first letters are the same and second and second last and so on. If there is an odd number the
programme will take that character as both. If the programme find any exceptions it returns false.

Problem 8

Write a function that merges two sorted lists into a new sorted list, e.g. merge([1,4,6], [2,3,5]) =
[1,2,3,4,5,6].



Problem 9

Write a function to calculate the square root of a number using Newton’s method. Newton’s method is to
approximate the square root of a number x by picking a starting point z and then repeating the following
operation.

After finding a method online to carry out newtons method and the square root function I commented the code. 
The for loop in thr main will call the methods using i from 1 - 10 and geting the square root of each of those
values comparing them and getting the differnece.

Problem 10

Write a function to reverse a string in Go.

I found a useful link on stackoverflow for a string reversal and added to it. Printing the string before and after
it had been reversed using for loops and rune for the unicode. 
