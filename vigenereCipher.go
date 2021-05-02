package vigenereCipher

import (
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
// The default is all upper and lowercase letters with numbers 0-9 and symbols
// Example input: "abcABC123"
func UpdateAlphabet(newAlphabet string) {
	alphabet = newAlphabet
	prepareAlphabet()
}

// Encrypt encrypts a message with the provided key
func Encrypt(text, key string) (string, error) {

	var enc string

	for i := 0; i < len(text); i++ {
		if numPos, ok := alphabetMap[string(text[i])]; ok {

			keyPos, err := getKeyPos(key, i)
			if err != nil {
				return "", err
			}

			x := (numPos + keyPos) % len(alphabet)

			enc += string(alphabet[x])

		} else {
			return "", fmt.Errorf("invalid plaintext character entered! %v", string(text[i]))
		}
	}

	return enc, nil
}

// Decrypt decrypts a message with the provided key
func Decrypt(enc, key string) (string, error) {

	var dec string

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

			dec += string(alphabet[x])

		} else {
			return "", fmt.Errorf("invalid plaintext character entered! %v", string(enc[i]))
		}
	}

	return dec, nil
}

func getKeyPos(key string, i int) (int, error) {

	keyPos, ok := alphabetMap[string(key[i%len(key)])]

	if ok {
		return keyPos, nil
	}
	return -1, fmt.Errorf("invalid key character entered! %v", string(string(key[i%len(key)])))
}

func prepareAlphabet() {
	alphabetMap = make(map[string]int)
	for i := 0; i < len(alphabet); i++ {
		alphabetMap[string(alphabet[i])] = i
	}
}
