package protocol

import (
	"fmt"
	"strconv"
)

// HashSize defines bitcoin hash size.
const HashSize = 32

// HashSize defines bitcoin hash size represented as string.
const HashStringSize = HashSize * 2

// Hash represents messaga data hashing output function.
type Hash [HashSize]byte

// NewHashFromString return Hash of 32 bytes from a string with max len of 64 chars.
func NewHashFromString(hash string) (*Hash, error) {
	if len(hash) > HashStringSize {
		return nil, fmt.Errorf("Invalid hash size (%d), size cannot be greather than (%d)", len(hash), HashStringSize)
	}

	val := []byte{}

	for i := 0; i < len(hash); i += 2 {
		byteValue, _ := strconv.ParseUint(hash[i:i+2], 16, 8)
		val = append(val, uint8(byteValue))
	}

	h := (*Hash)(val)

	return h, nil
}

// Reverse returns new Hash with the byte-wise reverse of computed hash.
func (h *Hash) Reverse() *Hash {
	reversedHash := []byte{}

	for i := len(h) - 2; i >= 0; i -= 2 {
		reversedHash = append(reversedHash, h[i:i+2]...)
	}

	return (*Hash)(reversedHash)
}
