package main

import (
	b64 "encoding/base64" // Aliasing a verbose immport.
	"fmt"
)

// Go provides built-in support for base64 encoding/decoding.

func main() {
	data := "abc123!?$*&()'-=@~" // The string we'll encode and decode.
	
	// Go supports standard and URL-compatible base64. Here's how to
	// encode using the standard encoder. The encoder rquires a []byte
	// so we can convert our string to that type.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoding may return an error, which you can check if you don't
	// already know if the input is well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64 format.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}