package main

import "fmt"

type person struct {
	name string
}

func main() {
	n := 10
	if n := ten(); n == 10 {
		fmt.Println("we got n=10 from branch scope")
		n = 11 // shadowed
	}
	fmt.Printf("n %d", n)
	tryToUpdateByValue()
	tryToUpdateByIndex()
	tryToUpdateByReference()
}

func tryToUpdateByValue() {
	people := []person{
		{name: "John"},
		{name: "Jane"},
		{name: "Amy"},
		{name: "Steve"},
	}

	for _, p := range people {
		fmt.Println(p.name)
		p.name = "Anonymous"
	}

	fmt.Println("[tryToUpdateByValue] checking if people names were updates")
	for _, p := range people {
		fmt.Println(p.name)
	}
}

func tryToUpdateByIndex() {
	people := []person{
		{name: "John"},
		{name: "Jane"},
		{name: "Amy"},
		{name: "Steve"},
	}

	for i, p := range people {
		fmt.Println(p.name)
		people[i].name = "Anonymous"
	}

	fmt.Println("[tryToUpdateByIndex] checking if people names were updates")
	for _, p := range people {
		fmt.Println(p.name)
	}
}

func tryToUpdateByReference() {
	people := []*person{
		{name: "John"},
		{name: "Jane"},
		{name: "Amy"},
		{name: "Steve"},
	}

	for _, p := range people {
		fmt.Println(p.name)
		p.name = "Anonymous"
	}

	fmt.Println("[tryToUpdateByReference] checking if people names were updates")
	for _, p := range people {
		fmt.Println(p.name)
	}
}

func ten() int {
	return 10
}
