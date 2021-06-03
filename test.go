package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

var c, python, java bool

const Pi = 3.14

var i, j int = 1, 2

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func swap(x, y string) (string, string) {
	return y, x
}
func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(add(42, 13))
	fmt.Println(math.Pi)
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	var f float64
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	var x, y int = 3, 4
	var z uint = uint(f)
	fmt.Println(x, y, z)
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)

}
