package chapters

import "fmt"

func slices() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]

	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)

	s := []int{1, 2, 3}
	s = append(s, 4)

	fmt.Println(s)

	s1 := s[1:]
	s2 := s[:2]
	s3 := s[1:2]

	fmt.Println(s1, s2, s3)
}
