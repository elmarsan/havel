package main

// SwapEndianUint32 swaps endianess order
// From big endian to little endian or vice versa
func SwapEndianUint32(b uint32) uint32 {
	b0 := b & 0x000000ff << 24
	b1 := b & 0x0000ff00 << 8
	b2 := b & 0x00ff0000 >> 8
	b3 := b & 0xff000000 >> 24

	swap := b0 | b1 | b2 | b3
	return swap
}
