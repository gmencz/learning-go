package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[:] {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}
}
