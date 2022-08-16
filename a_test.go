package a

import (
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func TestSecp256k1(t *testing.T) {
	_, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		t.Fatal(err)
	}
}
