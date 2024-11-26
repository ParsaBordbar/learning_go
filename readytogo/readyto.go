package main

import (
	"fmt"
    "readytogo/express" 
)

func main() {
	var kind string
	fmt.Scanf("%s", &kind)
	switch kind {
	case "express":
		express.Express()		
	}
}
