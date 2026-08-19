package main

import (
	"fmt"
	"os"
)

// penetration_testing_kit - Automated pentest toolkit
func penetration_testing_kit(path string) {
	fmt.Println("========================================")
	fmt.Println("  Penetration-Testing-Kit")
	fmt.Println("  Automated pentest toolkit")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	penetration_testing_kit(path)
}
