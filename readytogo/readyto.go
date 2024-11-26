package main

import (
	// "fmt"
	"os"
	// "os/exec"
)

func express () {
	dir_names := [5] string { "routes", "controllers", "middlewares", "models", "configs" }
	for i := 0 ; i < 5 ; i++ {
		if _, err := os.Stat(dir_names[i]); err != nil {
			os.Mkdir(dir_names[i], 0755)
		}
	}
	os.Chdir("routes")	
	if _, err := os.Stat("routes.js"); err != nil {
		os.Create("routes.js")
	}
}

func main () {
	express()
}