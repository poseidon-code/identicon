package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

/* Blocksize to Resolution mapping (using all bits in hash)
 * SHA256 : 64  = 8:8 or 1:1    (Square)
 * SHA512 : 128 = 16:8 or 2:1   (Rectangle)
 */

func SHA256(t string) (string, int, int) {
    return fmt.Sprintf("%x", sha256.Sum256([]byte(t))), 1, 1
}

func SHA512(t string) (string, int, int) {
    return fmt.Sprintf("%x", sha512.Sum512([]byte(t))), 2, 1
}
