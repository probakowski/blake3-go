package blake3_test

import (
	"blake3"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHasher(t *testing.T) {
	hasher := blake3.NewHasher()
	hasher.Write(make([]byte, 0))
	bytes := hasher.Finalize(64)
	fmt.Println(hex.EncodeToString(bytes))
}

func TestHash(t *testing.T) {
	fmt.Println(hex.EncodeToString(blake3.Hash(make([]byte, 0), 64)))
}
