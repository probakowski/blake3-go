package main

//#include "blake3.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"io"
	"unsafe"
)

type Blake3 interface {
	io.Writer
	Finalize(len int) []byte
}

type hasher struct {
	hash C.blake3_hasher
}

func (h *hasher) Write(p []byte) (n int, err error) {
	n = len(p)
	C.blake3_hasher_update(&h.hash, unsafe.Pointer(&p[0]), C.ulong(n))
	return
}

func (h *hasher) Finalize(len int) (out []byte) {
	out = make([]byte, len)
	C.blake3_hasher_finalize(&h.hash, (*C.uchar)(unsafe.Pointer(&out[0])), C.ulong(len))
	return
}

func NewHasher() Blake3 {
	hasher := hasher{}
	C.blake3_hasher_init(&hasher.hash)
	return &hasher
}

func main() {
	hasher := NewHasher()
	//fmt.Fprint(hasher, "Abcd")
	bytes := hasher.Finalize(64)
	fmt.Println(hex.EncodeToString(bytes))
}