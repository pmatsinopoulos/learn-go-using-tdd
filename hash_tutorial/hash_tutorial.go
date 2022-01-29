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

	privateKey := sha256.Sum256([]byte(input1))
	fmt.Printf("Size in bytes: %v\n", len(privateKey)) // should print "32" bytes length
	s = fmt.Sprintf("%x", privateKey)                  // hex representation
	fmt.Printf("%q has length: %d\n", s, len(s))       // should print "64" characters length
}
