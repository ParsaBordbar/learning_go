package main

import "fmt"

func prime_number (number int) (isPrime bool) {

	isPrime = true
	for i := 2 ; i <= number / 2 ; i++ {
		if number % i == 0 {
			isPrime = false
			break
		} 
	}  
	if isPrime {
		return
	}
	return
}

func main() {
	var number int;
	fmt.Scanf("%v" , &number)
	fmt.Println(prime_number(number))
}