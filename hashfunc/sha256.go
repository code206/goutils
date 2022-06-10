package hashfunc

import (
	"crypto/sha256"
	"encoding/hex"
)

func S2Sha256(s string) string {
	return B2Sha256([]byte(s))
}

func B2Sha256(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}
