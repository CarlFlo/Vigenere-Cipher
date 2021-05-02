# Vigenere Cipher

My implementation of the [Vigenere Cipher](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher) in [GO](https://golang.org/).

Test coverage: **93.5%**

## Features

- Encryption and decryption
- Supports custom alphabets

## Install

```
go get github.com/CarlFlo/vigenereCipher
```

## How to use

```go
// Encrypt a message with a key
output, err := vigenereCipher.Encrypt("Message", "Key")

// Decrypt a message with a key
output, err := vigenereCipher.Decrypt("Message", "Key")
```

## Options

You're able to change the alphabet by updating the alphabet string

```go
// This will update the code to use this alphabet
vigenereCipher.UpdateAlphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 ?!:;.,-_<>*'|=+{}()[]&#@/%\"")
```