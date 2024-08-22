package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// main Renames files in a directory by replacing spaces with dots and
// simplifying the resulting dots, and by uppercasing the first and last
// characters of a match to the regular expression 's[0-9]*e[0-9]*'.
//
// The directory path is provided as a command-line argument.
// No return values.
func main() {
	// Check if a directory path was provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <directory_path>\n", os.Args[0])
		os.Exit(1)
	}

	// Get the directory path from command-line argument
	dirPath := os.Args[1]

	// Read all files in the directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	sePattern := regexp.MustCompile(`(?i)s[0-9]+e[0-9]+`)

	// Iterate through each file
	for _, file := range files {
		if !file.IsDir() {
			oldName := file.Name()
			newName := strings.ReplaceAll(oldName, " ", ".")
			newName = simplifyDots(newName)

			newName = sePattern.ReplaceAllStringFunc(newName, func(match string) string {
				return strings.ToUpper(match[:1]) + strings.ToUpper(match[1:len(match)-1]) + strings.ToUpper(match[len(match)-1:])
			})

			// Construct full file paths
			oldPath := filepath.Join(dirPath, oldName)
			newPath := filepath.Join(dirPath, newName)

			// Rename the file
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Printf("Error renaming %s: %v\n", oldName, err)
			} else {
				fmt.Printf("Renamed: %s -> %s\n", oldName, newName)
			}
		}
	}
}

func simplifyDots(name string) string {
	for strings.Contains(name, "..") {
		name = strings.ReplaceAll(name, "..", ".")
	}
	return name
}
