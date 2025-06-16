package main

import "fmt"

func sum(a []int, target int) []int {

	m := make(map[int]int)

	for i, num := range a {
		b := target - num
		if j, c := m[b]; c {
			return []int{j, i}
		}
		m[num] = i
	}

	return nil
}

func main() {
	a := []int{2, 7, 11, 15}
	target := 9

	result := sum(a, target)

	if result != nil {
		fmt.Print("[", a[result[0]], ", ", a[result[1]], "] ")
		fmt.Print("(Because a[", result[0], "] + a[", result[1], "] == ", target, ")")
	} else {
		fmt.Println("there is no target found in the array")
	}

}
