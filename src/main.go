package main

import "fmt"
import "os"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: visualscriptTool.exe <path_to_visualscript_file>")
		return
	}
	filePath := os.Args[1]
	_, err := ReadVisualScriptFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading visual script: %v\n", err)
		return
	}
}
