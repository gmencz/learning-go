// Write a program that asks the user for a number n and gives them the possibility to choose
// between computing the sum and computing the product of 1,…,n.

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &n)

	fmt.Println("Awesome, now choose one of the following options:")
	fmt.Println("1: Compute the sum of 1 to", n)
	fmt.Println("2: Compute the product of 1 to", n)

	var option int
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}

		fmt.Println("The sum is:", sum)
	case 2:
		product := 1
		for i := 1; i <= n; i++ {
			product *= i
		}

		fmt.Println("The product is:", product)
	default:
		fmt.Println("Invalid option", option)
	}
}
