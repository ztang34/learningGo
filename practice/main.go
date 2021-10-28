package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	poodle := Dog{"poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Println(poodle.Breed)
	poodle.Weight = 9
	fmt.Println(poodle.Weight)
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}
