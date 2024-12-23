package main 

import "fmt"

func main () {
	// simple example
	basic_logger(summation(100, 400, 500, -90), "err 404", summation(500, -800, -600, 0))	
}

func summation (numbers ...int) (sum int) {
	for _, number := range numbers {
		sum += number
		fmt.Println(sum)
	}
	return
}

func basic_logger (logs ...interface{}) {
	for i, log := range logs {
		fmt.Printf("{ %d: %v }, ", i, log)
	}
}