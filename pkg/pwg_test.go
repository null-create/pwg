package pkg

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSecretGeneration(t *testing.T) {
	testSecret := GenSecret(64)

	assert.NotEqual(t, "", testSecret)
	assert.Equal(t, 64, len(testSecret))
}

func TestTextSecretGeneration(t *testing.T) {
	testSecret, err := GenPhraseSecret(4, true)
	if err != nil {
		t.Fatal(err)
	}

	words := strings.Split(testSecret, "-")

	assert.NotEqual(t, "", testSecret)
	assert.Equal(t, 4, len(words))
}
