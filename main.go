package main

import "awesomeProject/myfuncs/algs_sort"

func main() {

	arr := []float64{5, 1, -3, 3.2, 4.56, -12.34, 56, 12, -2.1, 100, 88, 82,
		76, 75, 66, 65, 64, 63, 59, 57, -45, -550, -999, -32455, -10356.45}

	algs_sort.Buble_sort(arr)
	algs_sort.Include_sort(arr)
	algs_sort.Fast_sort(&arr)

	//a := []float64{1, 2.5, -5}
	//b := []float64{-3.5, 2, 4.54}
	//fmt.Println("Скалярное произведение векторов: ")
	//fmt.Println(my_math.Mult_vector(a, b))
	//
	//matr1 := [][]float64{
	//	{1, 2},
	//	{3, 4}}
	//matr2 := [][]float64{
	//	{5, 6},
	//	{7, 8},
	//}
	//fmt.Println("\nПроизведение матриц: ")
	//fmt.Println(my_math.Mult_matrix(matr1, matr2))
}
