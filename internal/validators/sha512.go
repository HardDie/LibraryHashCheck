package validators

import (
	"bytes"
	"crypto/sha512"
)

type Sha512Validator struct {
	len int
}

func NewSha512() Sha512Validator {
	h := Sha512Validator{}
	h.len = calcHashLen(h)
	return h
}

func (v Sha512Validator) Name() string {
	return "sha512"
}

func (v Sha512Validator) Len() int {
	return v.len
}

func (v Sha512Validator) Hash(file []byte) []byte {
	hash := sha512.Sum512(file)
	return hash[0:]
}

func (v Sha512Validator) Validate(file, hash []byte) bool {
	fileHash := sha512.Sum512(file)
	return bytes.Equal(fileHash[0:], hash)
}
