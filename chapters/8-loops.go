package chapters

func loops() {

	// a simple for-loop
	for i := 0; i < 5; i++ {
		// this builded-in println() is for debugging only
		println(i)
	}

	// an infinite for-loop
	var j int
	for {
		println(j)
		j++
	}

	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	maps := map[string]int{"one": 1, "two": 2}
	// write-only variable
	for _, v := range maps {
		println(v)
	}
}
