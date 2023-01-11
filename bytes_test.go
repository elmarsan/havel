package main

import "testing"

func TestSwapEndianUint32(t *testing.T) {
	var b uint32 = 0xAABBCCDD
	var expected uint32 = 0xDDCCBBAA

	swap := SwapEndianUint32(b)
	if swap != expected {
		t.Errorf("Error while swaping endian. Expected: 0x%x, result: 0x%x", expected, swap)
	}
}
