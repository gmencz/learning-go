package main

import (
	"fmt"
	"math"
	"time"
)

// When two or more consecutive named function parameters share a type, you can omit
// the type from all but the last.
func add(x, y int) int {
	return x + y
}

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions, as with the example shown here.
// They can harm readability in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// The var statement declares a list of variables; as in function argument lists, the type is last.
// A var statement can be at package or function level.
var c, python, java bool

func main() {
	fmt.Println("The time is", time.Now())

	fmt.Println("Pi is", math.Pi)

	fmt.Println("2 + 3 is", add(2, 3))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var i int
	fmt.Println(i, c, python, java)

}

// NOTES

// Every Go program is made up of packages.
// Programs start running in package "main".

// It's considered a "good style" to group imports into a parenthesized, "factored" import statement:
// import (
// 	"fmt"
// 	"math"
// )

// In Go, a name is exported if it starts with a capital letter. For example, Pizza
// is an exported name, as is Pi, which is exported from the "math" package.

// Go's basic types are

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

// The expression T(v) converts the value v to the type T.
// i := 42
// f := float64(i)
// u := uint(f)
