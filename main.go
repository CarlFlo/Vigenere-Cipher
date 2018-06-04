package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Add any symbol here to add support for it
var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ?!:;.,-_<>*'|=+{}()[]&#@/%\""
var alphabetMap = make(map[string]int)

func main() {
	clear()

	startTime := time.Now().UTC()

	key, txt, _bool := handleArgs()

	createMap()

	if _bool { // Encrypt

		enc, err := encrypt(txt, key)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("Plaintext:", txt)
		fmt.Println("Key:", key, "\n")

		fmt.Print("Encrypted text between the > <\n>", enc, "<\n\n")

	} else if !_bool { // Decrypt

		dec, err := decrypt(txt, key)
		if err != nil {
			panic(err.Error())
		}

		fmt.Print("Decrypted text:\n", dec, "\n\n")

	}

	endTime := time.Now().UTC()
	duration := endTime.Sub(startTime)

	fmt.Println("Process took", duration.Nanoseconds()/1e6, "ms")
}

func handleArgs() (string, string, bool) {

	args := os.Args[1:]

	var key string
	var text string
	var encrypt bool // True for encryption. False for decryption

	if len(args) == 0 || args[0] == "help" {

		fmt.Println(fmt.Sprintf(`Must always have a key and choose only enc or dec
enc="Your string to encrypt"
dec="Your string to decrypt"
key="The key used to encrypt"
alp="abcdefghi..." if you want a custom alphabet

Default supported chars: 
%s

Example:
%s enc="Text to encrypt 123 :)" key="DH23C dSa"
%s dec="That to decrypt 123 :)" key="DH23C dSa"
`, alphabet, os.Args[0], os.Args[0]))

		os.Exit(3)
	}

	for i := 0; i < len(args); i++ {

		switch txt := (args[i])[:3]; txt {
		case "enc":
			text = (args[i])[4:]
			encrypt = true
		case "dec":
			text = (args[i])[4:]
			encrypt = false
		case "key":
			key = (args[i])[4:]
		case "alp": // For custom alphabet
			alphabet = (args[i])[4:]
		default:
			fmt.Println("Error! Invalid param:", txt)
			os.Exit(3)
		}
	}

	if len(key) == 0 || len(text) == 0 {
		fmt.Println("Args missing! ", args)
		os.Exit(3)
	}

	return key, text, encrypt
}

func encrypt(text, key string) (string, error) {

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

func decrypt(enc, key string) (string, error) {

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

func createMap() {

	for i := 0; i < len(alphabet); i++ {
		alphabetMap[string(alphabet[i])] = i
	}
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
