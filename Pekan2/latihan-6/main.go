package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCounty(address *Address) {
	address.Country = "Malaysia"
}

func main() {

	address1 := Address{"Subang", "Jabar", "Indonesia"}
	address2 := &address1
	address99 := Address{"Denpasar", "Bali", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(*address2)
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (value)   :", numberB)
	fmt.Println("numberB (address) :", numberB)

	ChangeCounty(&address99)
	fmt.Println(address99)

}
