package main

import (
	"fmt"
    "readytogo/express" 
	"readytogo/python"
)

func main() {
	var kind string
	fmt.Scanf("%s", &kind)
	switch kind {
	case "express":
		express.Express()
	case "python":
		python.Python()
	}
}
