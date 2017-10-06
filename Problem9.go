// Niall Devery
// Adapted from: https://gist.github.com/sighmin/9173219
package main

import (
    "fmt"
    "math"
)

func Newt(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        z = z - ((math.Pow(z, 2) - x) / (2 * z))// newtons square root method formula
    }
    return z
}

func Sqrt(x float64) float64 {
    return math.Sqrt(x)// square root function
}

func main() {
    times := 10
    for i := 0; i-1 < times; i++ {// incrementation of i and uses it for the equations
        sqrt := Sqrt(float64(i))// calls the square root function using i
        newt := Newt(float64(i))// calls the newtons method using i
        fmt.Println(i, "squared:")
        fmt.Println("  Sqrt:", sqrt)
        fmt.Println("  Newt:", newt)
        fmt.Println("  Difference:", math.Abs(sqrt-newt))
    }
}