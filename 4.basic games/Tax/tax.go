package main

import (
	"math"
	"fmt"
)

func main () {

    fmt.Println("Enter the amount of income: ")
    var income float64
    fmt.Scanf("%f", &income)
    var tax float64 = 0

	if income <= 100 {
		tax += income * 0.05
	} else if income > 100 && income <= 500 {
		tax += 100 * 0.05 + (income - 100) * 0.10
	} else if income > 500 && income <= 1000 {
		tax += 100 * 0.05 + 400 * 0.10 + (income - 500) * 0.15
	} else {
		tax += 100 * 0.05 + 400 * 0.10 + 500 * 0.15 + (income - 1000) * 0.20	
	}
	
	tax = math.Floor(tax)
	fmt.Println(tax)
}