package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	s := person{name:"Sean", age:50}
	fmt.Println(s.name)


	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
