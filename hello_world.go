package main

import (
	"fmt"
	"os"
	"strings"
)

//HelloWorld prints the text "Hello, world!" on the standard output if there are no command line arguments,
// otherwise it will print "Hello, <cmd-line-args>"
func HelloWorld() {
	if len(os.Args) > 1 {
		fmt.Printf("Hello, %s!\n", strings.Join(os.Args[1:], " "))
	} else {
		fmt.Println("Hello, world!")
	}
}
