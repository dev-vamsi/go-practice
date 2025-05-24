package main

import (
	"fmt"
	"os"
)

func SaveFileContent(filename string, content []byte) {
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func ReadFileContent(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []byte{}
	}

	return data
}
