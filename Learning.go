package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	var i, j int = 1, 2
	k := 3
	fmt.Println(i, j, k)
}

const Pi = 3.14

// Integer vs Float - Numeric constants
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

// For-loop
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// "while"-loop
func main() {
	sum := 1
	for sum < 1000 { // init-statement is dropped including semicolon
		sum += sum
	}
	fmt.Println(sum)
}

// if-statement
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if-else-statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}


// Square root function
package main

import (
	"fmt"
)

func Sqrt(x, err float64) float64 {
	z  := 1.0
	for (z*z-x) > err {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(5, 0.0000000001))
}


// Swift-Statement
// A switch statement is a shorter way to write a sequence of if - else statements. 
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// switch without condition -> Long if-then-else
func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// defer statement
// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 
// Deferred function calls are executed in last-in-first-out order. 

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}