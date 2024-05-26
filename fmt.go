package main

import "fmt"

func main() {
	firstName := "Eko"
	middleName := "Kurniawan"
	lastName := "Khannedy"

	fmt.Println("Hello '", firstName, middleName, lastName, "'")

	fmt.Printf("Hello '%s %s %s'\n", firstName, middleName, lastName)
}