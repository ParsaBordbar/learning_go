package python

import (
	"fmt"
	"log"
	"os"
)

func Python() {
	//get project name
	var proj_name string
	fmt.Print("Enter the Project name: ")
	fmt.Scanf("%s", &proj_name)

	if _, err := os.Stat(proj_name); os.IsNotExist(err) {
		if err := os.Mkdir(proj_name, 0755); err != nil {
			log.Fatalf("Failed to create project directory: %v", err)
		}
	}
	if err := os.Chdir(proj_name); err != nil {
		log.Fatalf("Failed to change directory to %s: %v", proj_name, err)
	}

	dir_and_files := map[string][]string{
		"src":   {"main.py"},
		"tests": {"test_main.py"},
		"data":  {"input.csv", "output.json"},
		"docs":  {"README.md", "LICENSE"},
	}

	// make directories and files
	for dir := range dir_and_files {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.Mkdir(dir, 0755); err != nil {
				log.Printf("Failed to create directory %s: %v", dir, err)
			}
		}

		if err := os.Chdir(dir); err != nil {
			log.Fatalf("Failed to change directory to %s: %v", proj_name, err)
		}

		for file := range dir_and_files[dir] {
			if _, err := os.Stat(dir_and_files[dir][file]); os.IsNotExist(err) {
				if _, err := os.Create(dir_and_files[dir][file]); err != nil {
					log.Printf("Failed to create file %v: %v", dir_and_files[dir][file], err)
				}
			}
		}

		if err := os.Chdir(".."); err != nil {
			log.Fatalf("Failed to change directory to %s: %v", proj_name, err)
		}
	}
}
