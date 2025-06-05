package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Usage: 
Windows: ./go-bootcamp.exe name-of-exercise
Linux/MacOS: ./go-bootcamp name-of-exercise`)
		return
	}

	name := os.Args[1]

	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	n := 1
	for _, f := range files {
		if f.IsDir() {
			oldName := f.Name()
			parts := strings.SplitN(oldName, "-", 2)
			if len(parts) >= 2 {
				i, err := strconv.Atoi(parts[0])
				if err == nil && i >= n {
					n = i + 1
				}
			}
		}
	}

	folder := fmt.Sprintf("%d-%s", n, name)
	if err := os.Mkdir(folder, 0755); err != nil {
		panic(err)
	}

	mainPath := filepath.Join(folder, "main.go")
	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Exercise: ` + name + `")
}
`
	if err := os.WriteFile(mainPath, []byte(mainContent), 0644); err != nil {
		panic(err)
	}

	cmd := exec.Command("code", mainPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println("Created:", folder)
}
