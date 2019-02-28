package main

func main() {
	foo := 2

	if foo == 1 {
		println("Bar")
	} else {
		println("Buz")
	}

	switch {
	case foo == 2:
		println("Case 2")
	case foo > 2:
		println("Case greater than 2")
	}

	switch foo2 := 2; foo2 {
	case 1:
		println("Case 1")
	case 2:
		println("Case 2")

	}

}
