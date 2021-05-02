# Vigenere Cipher

Supports lowercase, uppercase, numbers and spaces as well as some symbols.

Supports custom alphabets

## How to use

```go
// Encrypt a message with a key
output, err := vigenereCipher.Encrypt("Message", "Key")

// Decrypt a message with a key
output, err := vigenereCipher.Decrypt("Message", "Key")
```