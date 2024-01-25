package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to ByteText")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Printf("Editing file %s\n", filename)

	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}
	fmt.Printf("\nContents of %s:\n%s", filename, contents)

	fmt.Println("\nEnter new text below (type SAVE on a new line to save and exit)")
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "SAVE" {
			break
		}	
		lines = append(lines, line)
	}
	newContents := []byte{}

	for _, line := range lines {
		newContents = append(newContents, []byte(line)...)
		newContents = append(newContents, '\n')
	}

	err = os.WriteFile(filename, newContents, 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		os.Exit(1)
	}
	fmt.Println("File saves successfully")
}