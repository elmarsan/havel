package msg

import (
	"bytes"
	"testing"
)

func TestVarInDecode(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		len := []byte{0xaa}
		b := bytes.NewBuffer(len)

		varInt := &VarInt{}
		err := varInt.Decode(b)
		if err != nil {
			t.Errorf("Unable to decode uint8 (%s)", err.Error())
		}

		if varInt.Length != 170 {
			t.Errorf("Wrong uint8 decoding")
		}
	})

	t.Run("uint16", func(t *testing.T) {
		len := []byte{0xfd, 0x12, 0x1f}
		b := bytes.NewBuffer(len)

		varInt := &VarInt{}
		err := varInt.Decode(b)
		if err != nil {
			t.Errorf("Unable to decode uint16 (%s)", err.Error())
		}

		if varInt.Length != 7954 {
			t.Errorf("Wrong uint16 decoding")
		}
	})

	t.Run("uint32", func(t *testing.T) {
		len := []byte{0xfe, 0x12, 0x1f, 0x23, 0x22}
		b := bytes.NewBuffer(len)

		varInt := &VarInt{}
		err := varInt.Decode(b)
		if err != nil {
			t.Errorf("Unable to decode uint32 (%s)", err.Error())
		}

		if varInt.Length != 572727058 {
			t.Errorf("Wrong uint32 decoding")
		}
	})

	t.Run("uint64", func(t *testing.T) {
		len := []byte{0xff, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}
		b := bytes.NewBuffer(len)

		varInt := &VarInt{}
		err := varInt.Decode(b)
		if err != nil {
			t.Errorf("Unable to decode uint64 (%s)", err.Error())
		}

		if varInt.Length != 1229782938247303441 {
			t.Errorf("Wrong uint64 decoding")
		}
	})
}

func TestVarInEncode(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})

		varInt := &VarInt{
			Length: 170,
		}
		err := varInt.Encode(b)
		if err != nil {
			t.Errorf("Unable to encode uint8 (%s)", err.Error())
		}

		if b.Len() != 1 {
			t.Errorf("Wrong uint8 encoding")
		}
	})

	t.Run("uint16", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})

		varInt := &VarInt{
			Length: 7954,
		}
		err := varInt.Encode(b)
		if err != nil {
			t.Errorf("Unable to encode uint16 (%s)", err.Error())
		}

		if b.Len() != 2 {
			t.Errorf("Wrong uint16 encoding")
		}
	})

	t.Run("uint32", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})

		varInt := &VarInt{
			Length: 572727058,
		}
		err := varInt.Encode(b)
		if err != nil {
			t.Errorf("Unable to encode uint32 (%s)", err.Error())
		}

		if b.Len() != 4 {
			t.Errorf("Wrong uint32 encoding")
		}
	})

	t.Run("uint64", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})

		varInt := &VarInt{
			Length: 1229782938247303441,
		}
		err := varInt.Encode(b)
		if err != nil {
			t.Errorf("Unable to encode uint64 (%s)", err.Error())
		}

		if b.Len() != 8 {
			t.Errorf("Wrong uint64 encoding")
		}
	})
}
