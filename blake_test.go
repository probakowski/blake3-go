package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHasher(t *testing.T) {
	hasher := NewHasher()
	//fmt.Fprint(hasher, "Abcd")
	bytes := hasher.Finalize(64)
	fmt.Println(hex.EncodeToString(bytes))
}
