package hashfunc

import (
	"crypto/md5"
	"encoding/hex"
)

func S2Md5(s string) string {
	return B2Md5([]byte(s))
}

func B2Md5(b []byte) string {
	sum := md5.Sum(b)
	return hex.EncodeToString(sum[:])
}
