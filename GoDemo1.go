package main

import "fmt"

func main() {
	var (
		sal_emp1 int = 1000
		sal_emp2 int = 2000
		a int = 5
		b int = 3
	)
	sal_emp1 *= sal_emp2

	fmt.Printf("Total salary of employee 1 %d \n", sal_emp1)
	fmt.Print(sal_emp1 > sal_emp2, " \n")
	fmt.Println("sal_emp1> sal_emp2: ", sal_emp1 > sal_emp2)
	fmt.Println("sal_emp1< sal_emp2: ", sal_emp1 < sal_emp2)
	fmt.Println("sal_emp1== sal_emp2: ", sal_emp1 == sal_emp2)
	fmt.Println ("a & b = "a & b)
	fmt.Println ("a | b = "a | b)
	fmt.Println ("a ^ b = "a ^ b)
	fmt.Println ("a << b = "a << b)
	fmt.Println ("a >> b = "a >> b)
	
}
