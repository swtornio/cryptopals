package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Set 1, Challenge 1 - Convert hex to base64
// The string:
// 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
// Should produce:
// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
// So go ahead and make that happen. You'll need to use the Go programming language.
// No, we aren't going to break the rules of CTF all willy-nilly. Or are we?

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	decodedHex, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		return
	}

	encodedBase64 := base64.StdEncoding.EncodeToString(decodedHex)
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	fmt.Println("Encoded base64:", encodedBase64)

	if encodedBase64 == expectedOutput {
		fmt.Println("Success! The result is correct.")
	} else {
		fmt.Println("Failure. The result is incorrect.")
	}
}
