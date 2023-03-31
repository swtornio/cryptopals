package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

var englishFrequency = map[rune]float64{
	'a': 0.08167, 'b': 0.01492, 'c': 0.02782, 'd': 0.04253,
	'e': 0.12702, 'f': 0.02228, 'g': 0.02015, 'h': 0.06094,
	'i': 0.06966, 'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
	'm': 0.02406, 'n': 0.06749, 'o': 0.07507, 'p': 0.01929,
	'q': 0.00095, 'r': 0.05987, 's': 0.06327, 't': 0.09056,
	'u': 0.02758, 'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
	'y': 0.01974, 'z': 0.00074, ' ': 0.13000,
}

func xorSingleChar(input []byte, char byte) []byte {
	result := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i] ^ char
	}
	return result
}

func scoreEnglishText(input string) float64 {
	input = strings.ToLower(input)
	score := 0.0

	for _, ch := range input {
		if freq, ok := englishFrequency[ch]; ok {
			score += freq
		}
	}

	return score
}

func decryptSingleByteXOR(encoded []byte) (byte, string, float64) {
	var key byte
	var decrypted string
	maxScore := 0.0

	for char := byte(0); char <= 0x7F; char++ {
		xorResult := xorSingleChar(encoded, char)
		xorResultStr := string(xorResult)
		score := scoreEnglishText(xorResultStr)

		if score > maxScore {
			maxScore = score
			key = char
			decrypted = xorResultStr
		}
	}

	return key, decrypted, maxScore
}

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decodedInput, err := hex.DecodeString(input)
	if err != nil {
		fmt.Println("Error decoding hex:", err)
		return
	}

	key, decryptedText, _ := decryptSingleByteXOR(decodedInput)
	fmt.Printf("Decrypted text: %s\n", decryptedText)
	fmt.Printf("Key: 0x%02X ('%c')\n", key, key)
}
