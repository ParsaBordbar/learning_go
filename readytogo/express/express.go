package express

import (
	"fmt"
    "os"
    "log"
)

func Express() {

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


    dir_names := [5] string {"routes", "controllers", "middlewares", "models", "configs"}
	file_names := [2] string {"server.js", "app.js"}

    for _, dir := range dir_names {
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            if err := os.Mkdir(dir, 0755); err != nil {
                log.Printf("Failed to create directory %s: %v", dir, err)
            }
        }
    }

	for _, file := range file_names {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			if _, err := os.Create(file); err != nil {
				log.Printf("Failed to create file: %v", err)
			}
		}
	}

    if err := os.Chdir("routes"); err != nil {
        log.Fatalf("Failed to change directory to routes: %v", err)
    }
    if _, err := os.Stat("routes.js"); os.IsNotExist(err) {
        if _, err := os.Create("routes.js"); err != nil {
            log.Printf("Failed to create file routes.js: %v", err)
        }
    }
}
