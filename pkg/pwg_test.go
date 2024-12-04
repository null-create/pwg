package pkg

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestSecretGeneration(t *testing.T) {
	testSecret := GenSecret(64)

	assert.NotEqual(t, "", testSecret)
	assert.Equal(t, 64, len(testSecret))
}

func TestTextSecretGeneration(t *testing.T) {
	testSecret, err := GenPhraseSecret(4, false)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEqual(t, "", testSecret)
}
