package main

import (
	"fmt"
	"os"
	"os/exec"
)

func draw(MAP_X, MAP_Y, x, y  int) {
	// terminal := exec.Command("clear") 
	terminal := exec.Command("sl") 
	terminal.Stdout = os.Stdout
	terminal.Run()
	for i := 0 ; i < MAP_Y ; i++ {
		for j := 0 ; j < MAP_X ; j++ {
			if x == i && y == j {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {

	var MAP_X, MAP_Y, x, y, count int;
	fmt.Scanf("%v %v %v %v %v" , &MAP_X , &MAP_Y , &x , &y , &count)

	for i := 0 ; i < count ; i++ {
		var move int
		fmt.Scanf("%v" , &move)

		if move == 2 {
			y++
		} else if move == 4 {
			x--
		} else if move == 6 {
			x++
		} else if move == 8 {
			y--
		}
		draw(MAP_X, MAP_Y, x, y)
	}
}