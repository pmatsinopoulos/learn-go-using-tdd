package main

func Sum(arrayOfIntegers [4]int) (sum int) {
	sum = 0
	for _, v := range arrayOfIntegers {
		sum += v
	}
	return
}
