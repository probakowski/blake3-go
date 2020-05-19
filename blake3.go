package blake3

import (
	"io"
	"unsafe"
)

//#include "blake3.h"
//void hash(const void *input, size_t input_len, uint8_t *out, size_t out_len){
//    blake3_hasher hasher;
//    blake3_hasher_init(&hasher);
//    blake3_hasher_update(&hasher, input, input_len);
//    blake3_hasher_finalize(&hasher, out, out_len);
//}
import "C"

type Hasher interface {
	io.Writer
	Finalize(len int) []byte
}

type hasher struct {
	hash C.blake3_hasher
}

func (h *hasher) Write(p []byte) (n int, err error) {
	n = len(p)
	if n == 0 {
		p = make([]byte, 1)
	}
	C.blake3_hasher_update(&h.hash, unsafe.Pointer(&p[0]), C.size_t(n))
	return
}

func (h *hasher) Finalize(len int) (out []byte) {
	out = make([]byte, len)
	C.blake3_hasher_finalize(&h.hash, (*C.uchar)(unsafe.Pointer(&out[0])), C.size_t(len))
	return
}

func NewHasher() Hasher {
	hasher := hasher{}
	C.blake3_hasher_init(&hasher.hash)
	return &hasher
}

func Hash(data []byte, outputLength int) (out []byte) {
	out = make([]byte, outputLength)
	inputLength := len(data)
	if inputLength == 0 {
		data = make([]byte, 1)
	}
	C.hash(unsafe.Pointer(&data[0]), C.size_t(inputLength), (*C.uchar)(unsafe.Pointer(&out[0])), C.size_t(outputLength))
	return
}
