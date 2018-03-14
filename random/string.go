package random

import (
	"io"
	cryptorand "crypto/rand"
	mathrand "math/rand"
	"time"
	"fmt"
)

// returns a random string, length is length, and all characters are from set.
func NewString(length int, set string) string {
	bytes := []byte(set)
	result := []byte{}
	r := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	lenBytes := len(bytes)
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(lenBytes)])
	}
	return string(result)
}

// returns a random string, length is length. some characters may not be printed.
// also returns an error if something wrong.
func NewCryptoString(length int) (string, error) {
	r := make([]byte, length)
	_, err := io.ReadFull(cryptorand.Reader, r)
	return string(r), err
}

// Generate a random UUID according to RFC 4122
func NewUUID() (string, error) {
	len16Str, err := NewCryptoString(16)
	if err != nil {
		return "", err
	}
	uuid := []byte(len16Str)
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
