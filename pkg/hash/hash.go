package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

/* Blocksize to Resolution mapping (using all bits in hash)
 * SHA1   : 40  = 5:4           (not used)
 * SHA256 : 64  = 8:8 or 1:1    (Square)
 * SHA512 : 128 = 16:8 or 2:1   (Rectangle)
 */
func GenerateSHA256(t string) (hash string, w, h int) {
    w, h = 1, 1
    return fmt.Sprintf("%x", sha256.Sum256([]byte(t))), w, h
}

func GenerateSHA512(t string) (hash string, w, h int) {
    w, h = 2, 1
    return fmt.Sprintf("%x", sha512.Sum512([]byte(t))), w, h
}
