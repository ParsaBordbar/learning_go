package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// constants
const WIDTH = 20
const HEIGHT = 10

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

func draw(){
	
	ClearScreen()

	for j := 0 ; j <= HEIGHT ; j++ {

		// top and bottom walls
		if j == 0 || j == HEIGHT {
			for i := 0 ; i <= WIDTH ; i++ {
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
		for i := 0 ; i <= WIDTH ; i++ {

			if (i == 0  || i == WIDTH) && j != HEIGHT{
				fmt.Printf("#")
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Printf("\n")
	}
}

func main() {
	// game loop
	for {
		time.Sleep(time.Millisecond * 200)
		draw()
	}
}