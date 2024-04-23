package algs_sort

import "fmt"

func Partition(arr *[]float64, l int, r int) int {
	a := *arr
	x := a[r]
	less := l

	for i := l; i < r; i++ {
		if a[i] <= x {
			a[i], a[less] = a[less], a[i]
			less++
		}
	}
	a[less], a[r] = a[r], a[less]
	return less

}

func Fast_sortImpl(arr *[]float64, l int, r int) {
	if l < r {
		q := Partition(arr, l, r)
		Fast_sortImpl(arr, l, q-1)
		Fast_sortImpl(arr, q+1, r)
	}
}

func Fast_sort(arr *[]float64) {

	if arr != nil {
		Fast_sortImpl(arr, 0, len(*arr)-1)
	}

	fmt.Println(*arr)

}
