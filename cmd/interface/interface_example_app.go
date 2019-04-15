package main

import "fmt"

type Stringer interface {
	String() string
}

type person struct {
	fullName string
	age      int
}

func (p person) String() string {
	return p.fullName
}

func Print(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	persons := []person{
		{"Jane Doe", 18},
		{fullName: "Pete Ace", age: 19},
		{fullName: "Anne Caty", age: 22},
	}
	for i := 0; i < len(persons); i++ {
		Print(persons[i])
	}
}
