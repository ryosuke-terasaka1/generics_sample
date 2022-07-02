package test_code

type Numbers interface {
	~int | ~float64
}

func Average[T Numbers](s []T) T {
	var sum_num T
	for _, v := range s {
		sum_num += v
	}
	average := sum_num / T(len(s))
	return average
}
