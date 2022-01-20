package main

func Sum(arrayOfIntegers []int) (sum int) {
	sum = 0
	for _, v := range arrayOfIntegers {
		sum += v
	}
	return
}
