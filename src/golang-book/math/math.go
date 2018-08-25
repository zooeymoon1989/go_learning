package math

//这里有一个命名约定（name convention）如果是首字母大写的方法，则其他包或者程序能够看见它
func Average(xs []float64) (total float64) {

	total = float64(0)

	for _ , x := range xs {
		total += x
	}

	return

}
