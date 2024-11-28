package main

import (
	"flag"
	"fmt"
	"os"
	"readytogo/express"
	"readytogo/python"
)

func main() {
	// var kind string
	// fmt.Scanf("%s", &kind)

	pythonFlag := flag.NewFlagSet("python", flag.ExitOnError)
	pyProjectName := pythonFlag.String("name", "", "pythonProject")

	expressFlag := flag.NewFlagSet("express", flag.ExitOnError)
	exProjectName := expressFlag.String("name", "", "expressProject")

	if len(os.Args) < 2 {
		fmt.Println("Please enter an argument. [python, express, ...(Coming soon)]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "express":
		expressFlag.Parse(os.Args[2:])
		express.Express(*exProjectName)
	case "python":
		pythonFlag.Parse(os.Args[2:])
		python.Python(*pyProjectName)
	}
}
