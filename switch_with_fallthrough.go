package main

import (
	"fmt"
)

func main() {

	// The switch statement could also use a fallthrough keyword to enable execution
	// to follow to the next immediate case statement. In this case as long as
	// the case execution block has fallthrough as the last statement, execution
	// will follow to the next case statement
	switch "docker" {
	case "linux":
		fmt.Println("These are the linux courses ....")
	case "docker":
		fmt.Println("These are the available docker courses ...")
		fallthrough
	case "related":
		fmt.Println("These are some related courses ...")
	case "windows":
		fmt.Println("These are some windows courses ...")
	default:
		fmt.Println("There are no courses for what you have specified")
	}
}
