package main

import (
	"fmt"
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

	ch := make(chan string)
	go StreamPersonInfo(persons, ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}

func StreamPersonInfo(persons []person, ch chan string) {
	for _, p := range persons {
		var info string
		if p.likesCats {
			info = fmt.Sprintf("%s likes cats :)", p.fullName)
		} else {
			info = fmt.Sprintf("%s does not like cats :(", p.fullName)
		}
		ch <- info
	}
	close(ch)
}
