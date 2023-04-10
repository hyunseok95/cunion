

package lib

import (
	"crypto/rand"
	"io"
)

func tryReadRandom(p []byte) {
	io.ReadFull(rand.Reader, p)
}
