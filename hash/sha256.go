package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
