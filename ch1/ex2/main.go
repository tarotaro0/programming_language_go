package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("os.Args[%d]: %s\n", i, arg)
	}
}
