# Vigenere Cipher

My implementation of the [Vigenere Cipher](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher) in [GO](https://golang.org/).

This solution should obviously not be used to protect any real data.

Test coverage: **100.0%**

## Features

- Encryption and decryption with [polyalphabetic substitution](https://en.wikipedia.org/wiki/Polyalphabetic_cipher)
- Supports custom alphabets

## Install

```
go get github.com/CarlFlo/vigenereCipher
```

## How to use

```go
// Encrypt a message with a key
output, err := vigenereCipher.Encrypt("Message_to_encrypt", "Key")

// Decrypt a message with a key
output, err := vigenereCipher.Decrypt("Message_to_decrypt", "Key")
```

## Options

You're able to change the alphabet like this:

```go
// This is the default alphabet
vigenereCipher.UpdateAlphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ?!:;.,-_<>*'|=+{}()[]&#@/%\"")
```