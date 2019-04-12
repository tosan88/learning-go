package main

import (
	"fmt"
	"time"
)

type person struct {
	fullName  string
	likesCats bool
}

func main() {
	persons := []person{
		{"Jane Doe", false},
		{fullName: "Pete Ace"},
		{fullName: "Anne Caty", likesCats: true},
	}
	for _, p := range persons {
		go PrintPerson(p)
	}
	<-time.After(1 * time.Second)
}

//PrintPerson prints information about the given person p
func PrintPerson(p person) {
	if p.likesCats {
		fmt.Printf("%s likes cats :)\n", p.fullName)
	} else {
		fmt.Printf("%s does not like cats :(\n", p.fullName)
	}
}
