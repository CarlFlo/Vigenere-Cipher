# Vigenere Cipher

Supports lowercase, uppercase, numbers and spaces as well as some symbols.

Supports custom alphabets

Test coverage: **83.9%**

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