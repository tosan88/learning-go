package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Hello, %s!\n", strings.Join(os.Args[1:], " "))
	} else {
		fmt.Println("Hello, world!")
	}
}
