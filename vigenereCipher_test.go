package vigenereCipher

import "testing"

func TestEncrypt(t *testing.T) {

	input := "Example"
	key := "key"
	expectedOutput := "OBywtJo"

	output, err := Encrypt(input, key)

	if err != nil {
		t.Fatalf("Encryption failed %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("Encryption: Got '%s' expected '%s'", output, expectedOutput)
	}
}

func TestInvalidInput(t *testing.T) {

	// This character 'ä' is not in the default alphabet
	input := "ä"
	key := "key"

	_, err := Encrypt(input, key)

	if err == nil {
		t.Fatalf("Encryption should have failed")
	}
}

func TestDecrypt(t *testing.T) {

	input := "OBywtJo"
	key := "key"
	expectedOutput := "Example"

	output, err := Decrypt(input, key)

	if err != nil {
		t.Fatalf("Decryption failed %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("Decryption: Got '%s' expected '%s'", output, expectedOutput)
	}
}

func TestDecryptFail(t *testing.T) {

	input := "ä"
	key := "key"

	_, err := Decrypt(input, key)

	if err == nil {
		t.Fatalf("Decryption should have failed")
	}
}

func TestInvalidKey(t *testing.T) {

	input := "inputinputinput"
	key := "åäö"

	_, err := Encrypt(input, key)
	if err == nil {
		t.Fatalf("Encryption should have failed because of the key")
	}

	_, err = Decrypt(input, key)
	if err == nil {
		t.Fatalf("Decryption should have failed because of the key")
	}
}

func TestCustomAlphabet(t *testing.T) {

	UpdateAlphabet("ABC123ghj")
	input := "123"
	key := "A3h"
	expectedOutput := "1A1"

	output, err := Encrypt(input, key)

	if err != nil {
		t.Fatalf("Encryption failed %v", err)
	}

	if output != expectedOutput {
		t.Fatalf("Encryption: Got '%s' expected '%s'", output, expectedOutput)
	}

	output, err = Decrypt(output, key)

	if err != nil {
		t.Fatalf("Decryption failed %v", err)
	}

	if output != input {
		t.Fatalf("Decryption: Got '%s' expected '%s'", output, input)
	}
}
