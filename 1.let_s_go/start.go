package main

import "fmt"


func main () {

	fmt.Scanf("%v", &input)
	var gay_people = [4] string {"alireza", "Sadegh", "Ali", "Reza"}
	for i := 0 ; i < 4 ; i++ {
		fmt.Printf("value: %s, index: %d \n", gay_people[i], i)
	} 
}
