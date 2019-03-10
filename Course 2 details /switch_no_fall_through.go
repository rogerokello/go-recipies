package main

import (
	"fmt"
)

func main() {

	// The switch statement has an implicit break for each case block
	// so if a match is found, the code below the case is executed and
	// after execution continues to statements after the switch statement
	switch "docker" {
	case "linux":
		fmt.Println("These are the available linux courses .....")
	case "docker":
		fmt.Println("These are the available docker courses .....")
	case "windows":
		fmt.Println("These are the available windows courses .....")
	case "macosx":
		fmt.Println("These are the available macosx courses .....")
	default:
		fmt.Println("No course available for your selection")
	}
}
