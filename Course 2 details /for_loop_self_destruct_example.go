package main

import (
	"fmt"
	"time"
)

func main() {

	// A break statement may be used to break out of the loop
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)

	}
}
