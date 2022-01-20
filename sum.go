package main

func Sum(arrayOfIntegers [4]int) (sum int) {
	sum = 0
	for i := 0; i < 4; i++ {
		sum += arrayOfIntegers[i]
	}
	return
}
