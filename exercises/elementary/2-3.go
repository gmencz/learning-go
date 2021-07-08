// a) Write a program that asks the user for their name and greets them with their name.
// b) Modify the previous program such that only the users Alice and Bob are greeted with their names.

package main

import "fmt"

func main() {
	var name string

	fmt.Println("Enter your name:")
	fmt.Scanf("%s", &name)

	if name == "Alice" || name == "Bob" {
		fmt.Printf("Hey %s!\n", name)
	}
}
