package main

import (
	"awesomeProject/myfuncs"
	"awesomeProject/myfuncs/algs_sort"
	"fmt"
)

type User struct {
	UUID     int
	surname  string
	name     string
	patronym string
	age      int
}

type Student struct {
	OutUniversityData User
	code              string
	course            int
	group             int
	faculties         string
}

type Teacher struct {
	OutUniversityData User
	cafedra           string
	faculties         string
	science_rang      string
}

func f(value *int, x int) {
	*value = x
}

func ArrayToString(arr []int) string {
	res := ""
	for _, arg := range arr {
		res += string(arg) + " "
	}

	return res

}

func main() {
	//x := 4
	//value := &x

	//fmt.Print(x)
	//f(value, 12)
	//fmt.Print(x)

	//Изучим структуры данных

	arr := []int{
		1, 2, 3,
	}

	fmt.Println(arr)
	fmt.Println(ArrayToString(arr))
	myfuncs.Example()

	//for _, arg := range arr {
	//	fmt.Println(arg)
	//}

	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	algs_sort.Buble_sort([]float64{5, 1, -3, 3.2})

}
