package main

import (
	"fmt"
    "readytogo/express" 
	"readytogo/file_tree"
)

func main() {
	var kind string
	fmt.Scanf("%s", &kind)
	switch kind {
		case "express":
			express.Express()		
		case "tree": {
			root := "."
			fmt.Println("File Tree:")
			file_tree.PrintTree(root, "  ")
		}
	}
}
