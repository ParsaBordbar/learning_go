package main

import (
	"flag"
	"fmt"
	"os"
	"readytogo/express"
	"readytogo/python"
	"readytogo/file_tree"
)

func main() {

	pythonFlag := flag.NewFlagSet("python", flag.ExitOnError)
	pyProjectName := pythonFlag.String("name", "pythonProject", "name")

	expressFlag := flag.NewFlagSet("express", flag.ExitOnError)
	expressProjectName := expressFlag.String("name", "expressProject", "name")
	expressPrew := expressFlag.Bool("prew", false, "preview the file tree of an express app" )

	if len(os.Args) < 2 {
		fmt.Println("Please enter an argument. [python, express, ...(Coming soon)]")
		os.Exit(1)
	}

	switch os.Args[1] {
		case "express":{
			expressFlag.Parse(os.Args[2:])
			if *expressPrew {
				root := "." 
				fmt.Println("File Tree:")
				file_tree.PrintTree(root, "  ")
				os.Exit(0)
			} else {
				express.Express(*expressProjectName)
			}
		}
		case "python": {
			pythonFlag.Parse(os.Args[2:])
			python.Python(*pyProjectName)
		}
		// case "tree":
		// 	root := "."
		// 	fmt.Println("File Tree:")
		// 	file_tree.PrintTree(root, "  ")
	}

	fmt.Println(`
	██████╗ ███████╗ █████╗ ██████╗ ██╗   ██╗████████╗ ██████╗  ██████╗  ██████╗ 
	██╔══██╗██╔════╝██╔══██╗██╔══██╗╚██╗ ██╔╝╚══██╔══╝██╔═══██╗██╔════╝ ██╔═══██╗
	██████╔╝█████╗  ███████║██║  ██║ ╚████╔╝    ██║   ██║   ██║██║  ███╗██║   ██║
	██╔══██╗██╔══╝  ██╔══██║██║  ██║  ╚██╔╝     ██║   ██║   ██║██║   ██║██║   ██║
	██║  ██║███████╗██║  ██║██████╔╝   ██║      ██║   ╚██████╔╝╚██████╔╝╚██████╔╝
	╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═════╝    ╚═╝      ╚═╝    ╚═════╝  ╚═════╝  ╚═════╝ 
	` + "\n" + `	enjoy coding!` + "\n")
}
