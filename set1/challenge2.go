package main

import (
	"encoding/hex"
	"fmt"
)

// Set 1, Challenge 2 - Fixed XOR
// Write a function that takes two equal-length buffers and produces their XOR combination.
// If your function works properly, then when you feed it the string:
// 1c0111001f010100061a024b53535009181c
// ... after hex decoding, and when XOR'd against:
// 686974207468652062756c6c277320657965
// ... should produce:
// 746865206b696420646f6e277420706c6179

func xorBuffers(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("buffers have different lengths")
	}

	result := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}

	return result, nil
}

func main() {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	expectedOutput := "746865206b696420646f6e277420706c6179"

	decodedInput1, err := hex.DecodeString(input1)
	if err != nil {
		fmt.Println("Error decoding input1:", err)
		return
	}

	decodedInput2, err := hex.DecodeString(input2)
	if err != nil {
		fmt.Println("Error decoding input2:", err)
		return
	}

	xorResult, err := xorBuffers(decodedInput1, decodedInput2)
	if err != nil {
		fmt.Println("Error XOR'ing buffers:", err)
		return
	}

	encodedResult := hex.EncodeToString(xorResult)

	fmt.Println("Encoded XOR result:", encodedResult)

	if encodedResult == expectedOutput {
		fmt.Println("Success! The result is correct.")
	} else {
		fmt.Println("Failure. The result is incorrect.")
	}
}
