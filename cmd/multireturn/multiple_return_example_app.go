package main

import (
	"log"
	"os"
)

type person struct {
	fullName string
	age      int
}

func findPerson(persons []person, fullName string) (person, bool) {
	for _, p := range persons {
		if p.fullName == fullName {
			return p, true
		}
	}
	return person{}, false
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Too few arguments!")
	}
	persons := []person{
		{"Jane Doe", 18},
		{fullName: "Pete Ace", age: 19},
		{fullName: "Anne Caty", age: 22},
	}

	name := os.Args[1]
	p, found := findPerson(persons, name)
	if found {
		log.Printf("%s is in my list of persons. The age of this person is: %d", name, p.age)
	} else {
		log.Printf("%s is NOT in my list of persons.", name)
	}
}
