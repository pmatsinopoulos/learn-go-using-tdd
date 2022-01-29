package hash_tutorial

import (
	"crypto/sha256"
	"fmt"
)

func Tutorial() {
	const input1 = "Hello World!"

	h := sha256.New()
	h.Write([]byte(input1))
	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)
}
