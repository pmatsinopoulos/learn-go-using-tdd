package main

import "math"

type Circle struct {
	radius float64
}

func (circle Circle) Area() (area float64) {
	area = circle.radius * circle.radius * math.Pi
	return
}
