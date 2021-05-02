package vigenereCipher

import (
	"bytes"
	"fmt"
)

var (
	alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ?!:;.,-_<>*'|=+{}()[]&#@/%\""
	alphabetMap = make(map[string]int)
)

func init() {
	prepareAlphabet()
}

// UpdateAlphabet will change the alphabet the chipher
// uses for encryption and decryption.
// The default alphabet is: a-zA-Z0-9 with the addition of a few symbols.
//
// Example input: "abcABC123"
func UpdateAlphabet(newAlphabet string) {
	alphabet = newAlphabet
	prepareAlphabet()
}

// Encrypt encrypts a message with the provided key.
func Encrypt(text, key string) (string, error) {

	var buffer bytes.Buffer

	for i := 0; i < len(text); i++ {
		if numPos, ok := alphabetMap[string(text[i])]; ok {

			keyPos, err := getKeyPos(key, i)
			if err != nil {
				return "", err
			}

			x := (numPos + keyPos) % len(alphabet)

			buffer.WriteByte(alphabet[x])

		} else {
			return "", fmt.Errorf("invalid plaintext character entered! %v", string(text[i]))
		}
	}

	return buffer.String(), nil
}

// Decrypt decrypts a message with the provided key.
func Decrypt(enc, key string) (string, error) {

	var buffer bytes.Buffer

	for i := 0; i < len(enc); i++ {
		if numPos, ok := alphabetMap[string(enc[i])]; ok {
			keyPos, err := getKeyPos(key, i)
			if err != nil {
				return "", err
			}

			val := (numPos - keyPos)

			x := val % len(alphabet)

			if val < 0 {
				x = len(alphabet) + val
			}

			buffer.WriteByte(alphabet[x])

		} else {
			return "", fmt.Errorf("invalid plaintext character entered! %v", string(enc[i]))
		}
	}

	return buffer.String(), nil
}

func getKeyPos(key string, i int) (int, error) {

	char := string(key[i%len(key)])
	keyPos, ok := alphabetMap[char]

	if ok {
		return keyPos, nil
	}
	return -1, fmt.Errorf("invalid key character entered! %v", char)
}

func prepareAlphabet() {
	alphabetMap = make(map[string]int)
	for i := 0; i < len(alphabet); i++ {
		alphabetMap[string(alphabet[i])] = i
	}
}
