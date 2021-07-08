// Write a program that asks the user for a number n and prints the sum of the numbers 1 to n

package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println(sum)
}
