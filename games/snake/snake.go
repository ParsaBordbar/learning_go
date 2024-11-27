package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"time"
)

const WIDTH = 20
const HEIGHT = 10

var goingUp int = 0
var goingDown int = 0
var goingLeft int = 0
var goingRight int = 1

type Part struct {
	x int
	y int
}

var snake = []Part{
	{x: 7, y: 5},
	{x: 6, y: 5},
	{x: 5, y: 5},
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func draw() {
	ClearScreen()

	for j := 0; j <= HEIGHT; j++ {
		// top and bottom walls
		if j == 0 || j == HEIGHT {
			for i := 0; i <= WIDTH; i++ {
				if i == 0 {
					fmt.Printf(" ")
				} else if i == WIDTH {
					fmt.Printf(" \n")
				} else {
					fmt.Printf("-")
				}
			}
		}

		// main draw loop
		for i := 0; i <= WIDTH; i++ {
			if (i == 0 || i == WIDTH) && j != HEIGHT {
				fmt.Printf("#")
			} else {
				var something int = 0
				for _, part := range snake {
					if part.x == i && part.y == j {
						fmt.Printf("O")
						something = 1
						break
					}
				}

				if something == 0 {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func move() {
	// move the snake by adding a new head and removing the tail

	head := snake[0]
	var newHead Part = head

	// Determine the new head position
	if goingUp == 1 {
		newHead = Part{x: head.x, y: head.y - 1}
	} else if goingDown == 1 {
		newHead = Part{x: head.x, y: head.y + 1}
	} else if goingLeft == 1 {
		newHead = Part{x: head.x - 1, y: head.y}
	} else if goingRight == 1 {
		newHead = Part{x: head.x + 1, y: head.y}
	}

	// Check for collisions with walls or itself
	if newHead.x <= 0 || newHead.x >= WIDTH || newHead.y < 0 || newHead.y >= HEIGHT || slices.Contains(snake[1:], newHead) {
		fmt.Println("Game Over")
		os.Exit(0)
	}

	// Move the snake: append new head and remove tail
	snake = append([]Part{newHead}, snake[:len(snake)-1]...)
}

	func main() {
		// game loop
		for {
			move()      // Move the snake
			time.Sleep(time.Millisecond * 100)
			draw()      // Draw the updated game state
		}
	}
