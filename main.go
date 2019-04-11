package main

import "fmt"

func threeSum(a []int, t int) (int, int, int) {
	aInv := make(map[int]int)

	for i, Ai := range a {
		aInv[Ai] = i
	}
	// fmt.Println(aInv)

    n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			tPrime := t - a[i] - a[j]
			k, present := aInv[tPrime]
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
