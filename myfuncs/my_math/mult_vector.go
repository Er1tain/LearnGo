package my_math

func Mult_vector(a []float64, b []float64) float64 {

	var result float64 = 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			result += a[i] * b[i]
		}
	}
	return result

}
