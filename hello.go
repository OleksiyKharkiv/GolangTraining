package main

import "fmt"

func main() {
	var (
		companyName     string  = "GoTraining"
		yearEstablished int     = 2017
		name                    = "Oleks"
		myFloat         float32 = 3.14
		myInt           int     = 17
		myInt32         int32   = 787
		isActive        bool    = true
	)
	fmt.Println("Hello, World!")
	fmt.Println("Hello, Golang!")
	a, b, c, d := 1, 2, 3, 4
	fmt.Print(a, b, c, d)
	fmt.Println()
	fmt.Println(companyName, " ", yearEstablished)
	fmt.Printf("Hello %s, \n welcome to %s \n in year %d ! \n", name, companyName, yearEstablished)
	fmt.Printf("value: %s, type %T", companyName, companyName)
	fmt.Printf("value: %d, type %T", yearEstablished, yearEstablished)
	fmt.Printf("value: %s, type %T \n", name, name)
	fmt.Printf("value: %f, type %T \n", myFloat, myFloat)
	fmt.Printf("value: %d, type %T \n", myInt, myInt)
	fmt.Printf("Signet Integer (int): %d, type %T \n", myInt32, myInt32)
	fmt.Printf("value: %v, type %T \n", isActive, isActive)
}
