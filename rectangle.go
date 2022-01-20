package main

type Rectangle struct {
	height, width float64
}

func (rectangle Rectangle) Area() (area float64) {
	area = rectangle.height * rectangle.width
	return
}
