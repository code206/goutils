package hashfunc

import (
	"crypto/sha1"
	"encoding/hex"
)

func S2Sha1(s string) string {
	return B2Sha1([]byte(s))
}

func B2Sha1(b []byte) string {
	sum := sha1.Sum(b)
	return hex.EncodeToString(sum[:])
}
