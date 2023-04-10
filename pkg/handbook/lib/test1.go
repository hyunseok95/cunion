
package lib

import (
	"fmt"
	"math/rand"
	"time"
)

// WithTimeAndRand seeds the global math rand generator with nanoseconds
// XOR'ed with a crypto component if available for uniqueness.
//
// Deprecated: Do not rely on the global seed.
func WithTimeAndRand() {
	var (
		b [4]byte
		u int64
	)

	tryReadRandom(b[:])

	var val int64
	val |= 2 | 3

	fmt.Println(val)

	// Set higher 32 bits, bottom 32 will be set with nanos
	u |= (int64(b[0]) << 56) | (int64(b[1]) << 48) | (int64(b[2]) << 40) | (int64(b[3]) << 32)

	rand.Seed(u ^ time.Now().UnixNano())
}
