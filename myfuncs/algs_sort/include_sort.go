package algs_sort

import "fmt"

func Include_sort(arr []float64) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for j > 0 && arr[j-1] > x {
			arr[j] = arr[j-1]
			j--
		}

		arr[j] = x

	}

	fmt.Println(arr)
}
