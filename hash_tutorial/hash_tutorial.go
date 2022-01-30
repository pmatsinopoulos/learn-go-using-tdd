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
	fmt.Printf("Message digiest of %q is %q\n", input1, s)

	messageDigest := sha256.Sum256([]byte(input1))
	fmt.Printf("Size in bytes of message digest: %v\n", len(messageDigest)) // should print "32" bytes length
	s = fmt.Sprintf("%x", messageDigest)                                    // hex representation
	fmt.Printf("%q has length: %d\n", s, len(s))                            // should print "64" characters length

	fmt.Println("Generating a private key...")
	privateKey := sha256.New().Sum([]byte("my-password")) // private keys are a representation of a password "my-password"
	privateKeyStr := fmt.Sprintf("%x", privateKey)
	fmt.Printf("Len in bytes of private key: %d\n", len(privateKey))
	fmt.Printf("Private key in hex format: %q\n", privateKeyStr)
	fmt.Printf("Length of private key in hex representation: %d\n", len(privateKeyStr))

}
