// Write a program that prints a multiplication table for numbers up to 12.

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &n)

	for i := 1; i <= 12; i++ {
		result := n * i
		fmt.Printf("%dx%d = %d\n", n, i, result)
	}
}
