package main

import "fmt"

func main_() {
	x := [3]string{"foo", "bar", "mary"}

	y := x[:] // slice on array x

	z := make([]string, len(x)) // create slice with capacity equal to the capacity of +x+
	copy(z, x[:])               // now slice +z+ is backed by an array which has values equal to the values of +x+,
	// but they are different

	y[1] = "Panos"
	fmt.Printf("x = %v\n", x) // expect "foo", "Panos", "mary"

	fmt.Printf("y = %v\n", y) // expect "foo", "Panos", "mary"

	fmt.Printf("z = %v\n", z) // expect "foo", "bar", "mary"

}
