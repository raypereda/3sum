package main

import "fmt"

// i + j + k = x
// i + j - x = -k
// -(i + j - x) = k


func threeSum(a []int, x int) (int, int, int) {
	invA := make(map[int]int)
    for k, kValue := range a {
		invA[kValue] = k
	}

	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			// does this exist in a, with different i and j
			kValue := -(a[i] + a[j] - x)
			k, ok := invA[kValue]
			if ok && k != i && k != j {
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