package python

import (
	"fmt"
	"log"
	"os"
)

func Python(proj_name string) {

    // To implement: change output to desired route in next update
    outputDir := "output"

    if _, err := os.Stat(outputDir); os.IsNotExist(err) {
        if err := os.Mkdir(outputDir, 0755); err != nil {
            log.Fatalf("Failed to create project directory: %v", err)
        }
    }
    if err := os.Chdir(outputDir); err != nil {
        log.Fatalf("Failed to change directory to %s: %v", outputDir, err)
    }

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
		".":     {"requirements.txt", "setup.py"},
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
			if _, err := os.Stat(dir_and_files[dir][file]); err != nil {
				os.Create(dir_and_files[dir][file])

				switch dir_and_files[dir][file] {
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

		if err := os.Chdir(".."); err != nil {
			log.Fatalf("Failed to change directory to %s: %v", proj_name, err)
		}
	}
}
