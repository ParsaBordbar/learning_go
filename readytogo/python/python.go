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

	file_names := [2]string{"main.py", "requirements.txt"}
	dir_names := []string{"src", "tests", "data", "docs"}

	// make directories
	for _, dir := range dir_names {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.Mkdir(dir, 0755); err != nil {
				log.Printf("Failed to create directory %s: %v", dir, err)
			}
		}
	}

	// make files
	for i := 0; i < 2; i++ {
		if _, err := os.Stat(file_names[i]); err != nil {
			os.Create(file_names[i])

			switch file_names[i] {
			case "main.py":
				// Append string to file
				f, _ := os.OpenFile("main.py", os.O_APPEND|os.O_WRONLY, 0644)
				init_code := []string{"if __name__ == \"__main__\":", "\tpass"}
				for _, v := range init_code {
					fmt.Fprintln(f, v)
				}
			}
		}
	}
}
