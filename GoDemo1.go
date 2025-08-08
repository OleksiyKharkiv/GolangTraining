package main

import "fmt"

func main() {
	var (
		sal_emp1 int = 1000
		sal_emp2 int = 2000
	)
	sal_emp1 *= sal_emp2

	fmt.Printf("Total salary of employee 1 %d \n", sal_emp1)
}
