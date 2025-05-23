package pkg

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

const chars string = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!?@#$%^&*()_+={}:;><,.'"|\/`

// info about secret generation: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

// generate a random string of n length
func GenSecret(length int) string {
	charLen := len(chars)
	b := make([]byte, length)
	// generates len(b) random bytes.
	// bytes serve as indicies into the chars string
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("failed to generate secret: %v", err)
	}
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%charLen]
	}
	return string(b)
}

// get random word based secret, with optional dashes between words
func GenPhraseSecret(totalWords int, useDashes bool) (string, error) {
	var secret string
	for i := 0; i < totalWords; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(bundledDictionary))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random index: %w", err)
		}
		word := bundledDictionary[index.Int64()]
		if useDashes {
			secret += fmt.Sprintf("%s-", word)
		} else {
			secret += word
		}
	}
	if useDashes {
		secret = secret[:len(secret)-1] // remove the last dash
	}
	return secret, nil
}
