package main

import "fmt"

func main() {
	var (
		companyName     string = "GoTraining"
		yearEstablished int    = 2017
		name                   = "Oleks"
	)
	fmt.Println("Hello, World!")
	fmt.Println("Hello, Golang!")
	a, b, c, d := 1, 2, 3, 4
	fmt.Print(a, b, c, d)
	fmt.Println()
	fmt.Println(companyName, " ", yearEstablished)
	fmt.Printf("Hello %s, \n welcome to %s \n in year %d ! \n", name, companyName, yearEstablished)
	fmt.Printf("value: %s, type %T", companyName, companyName)
}
