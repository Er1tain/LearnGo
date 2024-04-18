package algs_sort

import "fmt"

func Buble_sort(arr []float64) {

	for i := 0; i+1 < len(arr); i++ {
		for j := 0; j+1 < len(arr)-i; j++ {

			if arr[j+1] < arr[j] {
				c := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = c
			}

		}
	}

	fmt.Println(arr)

}
