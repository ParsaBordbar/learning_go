package express

import (
	// "fmt"
    "os"
    "log"
    "os/exec"
)

// To implement: enums and arguments --flags for pure & ts
func Express(proj_name string) {
    
    // type flag string
    // const (
    //     pure    Flag = "pure"
    //     ts      Flag = "ts"
    // )

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
    
    // Create project directory if not exists
    if _, err := os.Stat(proj_name); os.IsNotExist(err) {
        if err := os.Mkdir(proj_name, 0755); err != nil {
            log.Fatalf("Failed to create project directory: %v", err)
        }
    }
    if err := os.Chdir(proj_name); err != nil {
        log.Fatalf("Failed to change directory to %s: %v", proj_name, err)
    }

    // npm package init
    terminal := exec.Command("npm", "init", "-y") 
    terminal.Stdout = os.Stdout
    terminal.Run()

    terminal2 := exec.Command("npm", "install", "express", "nodemon", "joi", "mongoose", "helmet", "morgan", "axios")
    terminal2.Stdout = os.Stdout
    terminal2.Stderr = os.Stderr
    err := terminal2.Run()
    if err != nil {
        log.Printf("Error installing packages: %v\n", err)
    }

    // Create Sub-dir and files if not exist
    dir_names := [5] string {"routes", "controllers", "middlewares", "models", "configs"}
	file_names := [2] string {"index.js", "app.js"}

    // Create sub-dir
    for _, dir := range dir_names {
        if _, err := os.Stat(dir); os.IsNotExist(err) {
            if err := os.Mkdir(dir, 0755); err != nil {
                log.Printf("Failed to create directory %s: %v", dir, err)
            }
        }
    }
    // Create Files
	for _, file := range file_names {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			if _, err := os.Create(file); err != nil {
				log.Printf("Failed to create file: %v", err)
			}
		}
	}

    // Go to dir routes and create routes.js if not exist
    if err := os.Chdir("routes"); err != nil {
        log.Fatalf("Failed to change directory to routes: %v", err)
    }
    if _, err := os.Stat("routes.js"); os.IsNotExist(err) {
        if _, err := os.Create("routes.js"); err != nil {
            log.Printf("Failed to create file routes.js: %v", err)
        }
    }
}
