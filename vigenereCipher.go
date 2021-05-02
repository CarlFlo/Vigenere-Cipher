package vigenereCipher

import (
	"fmt"
)

var (
	alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ?!:;.,-_<>*'|=+{}()[]&#@/%\""
	alphabetMap = make(map[string]int)
)

func init() {
	for i := 0; i < len(alphabet); i++ {
		alphabetMap[string(alphabet[i])] = i
	}
}

func Encrypt(text, key string) (string, error) {

	var enc string

	for i := 0; i < len(text); i++ {
		if numPos, ok := alphabetMap[string(text[i])]; ok {
			var keyPos int
			if _keyPos, _ok := alphabetMap[string(key[i%len(key)])]; _ok {
				keyPos = _keyPos
			} else {
				return "", fmt.Errorf("Invalid key character entered! %v", string(string(key[i%len(key)])))
			}

			x := (numPos + keyPos) % len(alphabet)

			enc += string(alphabet[x])

		} else {
			return "", fmt.Errorf("Invalid plaintext character entered! %v", string(text[i]))
		}
	}

	return enc, nil
}

func Decrypt(enc, key string) (string, error) {

	var dec string

	for i := 0; i < len(enc); i++ {
		if numPos, ok := alphabetMap[string(enc[i])]; ok {
			var keyPos int
			if _keyPos, _ok := alphabetMap[string(key[i%len(key)])]; _ok {
				keyPos = _keyPos
			} else {
				return "", fmt.Errorf("Invalid key character entered! %v", string(string(key[i%len(key)])))
			}

			val := (numPos - keyPos)

			x := val % len(alphabet)

			if val < 0 {
				x = len(alphabet) + val
			}

			dec += string(alphabet[x])

		} else {
			return "", fmt.Errorf("Invalid plaintext character entered! %v", string(enc[i]))
		}
	}

	return dec, nil
}
