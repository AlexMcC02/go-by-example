package main

import (
	"crypto/sha256" // Go implements several hash functions in various crypto/* packages.
	"fmt"
)

// SHA256 hashes are frequently used to compute short identities for
// binary or text blobs. For example, TLS/SSL certificates use SHA256
// to compute a certificate's signature.

func main() {
	s := "The lands of colchis draw near..."
	h := sha256.New() // New hash instance of SHA256.
	h.Write([]byte(s)) // The Write function expects a bytes, here we use type coercion.
	bs := h.Sum(nil) // This obtains the finalised hash result as a byte slice.

	fmt.Println(s) // Printing the original string.
	fmt.Printf("%x\n", bs) // Printing the hashed output.
}