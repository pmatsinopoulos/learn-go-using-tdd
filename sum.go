package main

func Sum(arrayOfIntegers []int) (sum int) {
	sum = 0
	for _, v := range arrayOfIntegers {
		sum += v
	}
	return
}

func SumAll(slices ...[]int) (result []int) {
	result = make([]int, len(slices))
	for i, s := range slices {
		result[i] = Sum(s)
	}
	return
}

func SumAllTails(slices ...[]int) (result []int) {
	result = make([]int, len(slices))
	for index, slice := range slices {
		if len(slice) >= 1 {
			result[index] = Sum(slice[1:])
		}
	}
	return
}
