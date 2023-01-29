package msg

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// VarInt represents variable length integer using in messaging.
// https://en.bitcoin.it/wiki/Protocol_documentation#Variable_length_integer
// < 0xfd - uint8 reading 1 byte
// <= 0xffff - uint16 reading 3 bytes
// <= 0xffff ffff - uint32 reading 5 bytes
// uint64 reading 9 bytes
type VarInt struct {
	Length uint
}

// Decode decodes VarInt from r.
func (vi *VarInt) Decode(r io.Reader) error {
	var b1 uint8

	err := Decode(r, binary.LittleEndian, &b1)
	if err != nil {
		return err
	}

	if b1 < 0xfd {
		vi.Length = uint(b1)
		return nil
	}

	suffix := b1 & 0x0f
	switch suffix {
	case 0xd:
		{
			var len uint16
			err := Decode(r, binary.LittleEndian, &len)
			if err != nil {
				return err
			}

			vi.Length = uint(len)
			return nil
		}
	case 0xe:
		{
			var len uint32
			err := Decode(r, binary.LittleEndian, &len)
			if err != nil {
				return err
			}

			vi.Length = uint(len)
			return nil
		}
	case 0xf:
		{
			var len uint64
			err := Decode(r, binary.LittleEndian, &len)
			if err != nil {
				return err
			}

			vi.Length = uint(len)
			return nil
		}
	}

	return fmt.Errorf("Wrong var int")
}

// Encode  encodes VarInt into w.
func (vi *VarInt) Encode(w io.Writer) error {
	len := vi.Length

	switch {
	case len <= math.MaxUint8:
		{
			b := uint8(len)
			err := Encode(w, binary.LittleEndian, &b)
			if err != nil {
				return err
			}
		}
	case len <= math.MaxUint16:
		{
			b := uint16(len)
			err := Encode(w, binary.LittleEndian, &b)
			if err != nil {
				return err
			}
		}
	case len <= math.MaxUint32:
		{
			b := uint32(len)
			err := Encode(w, binary.LittleEndian, &b)
			if err != nil {
				return err
			}
		}
	default:
		{
			b := uint64(len)
			err := Encode(w, binary.LittleEndian, &b)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
