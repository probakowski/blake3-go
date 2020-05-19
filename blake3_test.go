package blake3

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHasher(t *testing.T) {
	hasher := NewHasher()
	hasher.Write(make([]byte, 0))
	bytes := hasher.Finalize(64)
	fmt.Println(hex.EncodeToString(bytes))
}

func TestHash(t *testing.T) {
	fmt.Println(hex.EncodeToString(Hash(make([]byte, 0), 64)))
}
