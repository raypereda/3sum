package main

import "fmt"

func threeSum(a []int, x int) (int, int, int) {
	m := make(map[int]int)

	for k, aK := range a {
		xPrime := -aK + x 
		m[xPrime] = k
	}
	// fmt.Println(m)

	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			tPrime := a[i] + a[j]
			k, present := m[tPrime]
			// fmt.Println(i,j,k)
			if present && k != i && k != j {
				return a[i], a[j], a[k]
			}

		}
	}
	return -1, -1, -1
}

func main() {
	a := []int{3, 1, 2, 4}

	fmt.Println(threeSum(a, 7))
}
