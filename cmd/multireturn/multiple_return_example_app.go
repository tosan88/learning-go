package main

import "fmt"

type person struct {
	fullName  string
	likesCats bool
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
	persons := []person{
		{"Jane Doe", false},
		{fullName: "Pete Ace"},
		{fullName: "Anne Caty", likesCats: true},
	}
	p, found := findPerson(persons, "Jane Doe")
	if found {
		fmt.Printf("%s is in my list of persons.", p.fullName)
	} else {
		fmt.Printf("%s is NOT in my list of persons.", p.fullName)
	}
}
