package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if the target directory is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <target_directory> [exceptions...]")
		return
	}

	// Get the target directory from the first argument
	targetDir := os.Args[1]

	// Get the list of exceptions from the remaining arguments
	exceptions := os.Args[2:]

	// Build a map for faster lookup of exceptions
	exceptionsMap := make(map[string]bool)
	for _, ex := range exceptions {
		exceptionsMap[ex] = true
	}

	// Check if the target directory exists and is a directory
	info, err := os.Stat(targetDir)
	if err != nil {
		fmt.Printf("Error accessing target directory: %v\n", err)
		return
	}
	if !info.IsDir() {
		fmt.Printf("Target path is not a directory: %s\n", targetDir)
		return
	}

	// Print the root folder icon
	fmt.Println("📁")

	// Start printing the directory tree
	printDirectoryTree(targetDir, "", "", exceptionsMap)
}

func printDirectoryTree(rootPath, relativePath, prefix string, exceptions map[string]bool) {
	currentPath := filepath.Join(rootPath, relativePath)

	// Read the directory entries, skip if error
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		fmt.Printf("Error reading directory '%s': %v\n", currentPath, err)
		return
	}

	// Filter entries, only directories and not in exceptions
	var dirs []os.DirEntry
	for _, entry := range entries {
		if entry.IsDir() {
			if exceptions[entry.Name()] {
				continue
			}
			dirs = append(dirs, entry)
		}
	}

	// Iterate over directories
	for i, entry := range dirs {
		isLast := i == len(dirs)-1

		// Choose branch character
		var branch string
		if isLast {
			branch = "└──📁 "
		} else {
			branch = "├──📁 "
		}

		// Print the directory name with the prefix and branch
		fmt.Printf("%s%s%s\n", prefix, branch, entry.Name())

		// Build new prefix for the next level
		var newPrefix string
		if isLast {
			newPrefix = prefix + "    "
		} else {
			newPrefix = prefix + "│   "
		}

		// Recursively print subdirectories
		nextRelPath := filepath.Join(relativePath, entry.Name())
		printDirectoryTree(rootPath, nextRelPath, newPrefix, exceptions)
	}
}
