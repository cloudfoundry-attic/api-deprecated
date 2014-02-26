package digest

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func Hex(r io.Reader) (digest string, err error) {
	hasher := sha1.New()
	_, err = io.Copy(hasher, r)
	if err != nil {
		return
	}
	digest = hex.EncodeToString(hasher.Sum(nil))
	return
}
