package hash_tutorial

import (
	"crypto/sha256"
	"encoding/hex"
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

func CalculateMerkleRootExample() {
	hA := "51d37bdd871c9e1f4d5541be67a6ab625e32028744d7d4609d0c37747b40cd2d"
	hB := "60c25dda8d41f8d3d7d5c6249e2ea1b05a25bf7ae2ad6d904b512b31f997e1a1"
	//hC := "01f314cdd8566d3e5dbdd97de2d9fbfbfd6873e916a00d48758282cbb81a45b9"
	//hD := "b519286a1040da6ad83c783eb2872659eaf57b1bec088e614776ffe7dc8f6d01"

	hAbytes, _ := hex.DecodeString(hA)
	hBbytes, _ := hex.DecodeString(hB)
	hAplusBbytes := append(hAbytes, hBbytes...)
	//hAplusBbytes, _ := hex.DecodeString(fmt.Sprintf("%s%s", hA, hB))

	hAplusB := sha256.Sum256(hAplusBbytes)
	hAplusBStr := fmt.Sprintf("%x", hAplusB)
	fmt.Printf("...first hash A+B = %q\n", hAplusBStr)
	hAplusBbytes, _ = hex.DecodeString(hAplusBStr)
	hAplusB = sha256.Sum256(hAplusBbytes)
	hAplusBStr = fmt.Sprintf("%x", hAplusB)
	fmt.Printf("H for A+B = %q\n", hAplusBStr)
}
