package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(
		"The user name currently logged into your os is ",
		os.Getenv("USERNAME"))
}
