package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

	// Analogous to while in other languages
	i := 0
	for {
		i++
		println(i)
		if i > 5 {
			break
		}
	}

	// Looping over arrays, slices and maps

	// Looping over arrays
	s := []string{"one", "two", "three"}

	for idx, v := range s {
		println(idx, v)
	}

	// Looping over maps
	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "bar"
	m["third"] = "buz"

	for k, v := range m {
		println(k, v)
	}

}
