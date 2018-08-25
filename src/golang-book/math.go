package math

func Average(xs []float64) (total float64) {

	total = float64(0)

	for _ , x := range xs {
		total += x
	}

	return

}
