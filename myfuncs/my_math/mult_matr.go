package my_math

func Mult_matrix(matr1 [][]float64, matr2 [][]float64) [][]float64 {

	//initial result matrix zero
	result := make([][]float64, len(matr1))
	for i := 0; i < len(result); i++ {
		result[i] = make([]float64, len(matr1))
		for j := 0; j < len(matr1); j++ {
			result[i][j] = 0
		}
	}

	if len(matr1) == len(matr2) {
		n := len(matr1)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					result[i][j] += matr1[i][k] * matr2[k][j]
				}
			}
		}
	}

	return result
}
