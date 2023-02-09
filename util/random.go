package util

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func RandomInt(min, max int64) int64{
	return min + math_rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i< n; i++ {
		c := alphabet[math_rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 100)
}

func RandomCurrency() string {
	currencies := []string {"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[math_rand.Intn(n)]
}