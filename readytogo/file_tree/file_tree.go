package file_tree

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintTree(root string, indent string) {
	entries, err := os.ReadDir(root)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(indent + "|-- " + entry.Name())
		if entry.IsDir() {
			PrintTree(filepath.Join(root, entry.Name()), indent+"    ")
		}
	}
}
