package msg

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestReadUint8(t *testing.T) {
	d := []byte{0xaa}
	b := bytes.NewBuffer(d)

	t.Run("should read uint8", func(t *testing.T) {
		var val uint8
		err := readUint8(b, binary.LittleEndian, &val)
		if err != nil {
			t.Error("Could not read uint8")
		}

		if val != 0xaa {
			t.Errorf("Wrong uint8 read: actual (0x%x), expected (0xaa)", val)
		}
	})

	t.Run("should NOT read uint8", func(t *testing.T) {
		b.Next(1)

		var val uint8
		err := readUint8(b, binary.LittleEndian, &val)
		if err == nil {
			t.Error("An error should have occurred reading uint8")
		}
	})
}

func TestReadUint16(t *testing.T) {
	d := []byte{0xaa, 0xbb}
	b := bytes.NewBuffer(d)

	t.Run("should read uint16", func(t *testing.T) {
		var val uint16
		err := readUint16(b, binary.LittleEndian, &val)
		if err != nil {
			t.Error("Could not read uint16")
		}

		if val != 0xbbaa {
			t.Errorf("Wrong uint16 read: actual (0x%x), expected (0xbbaa)", val)
		}
	})

	t.Run("should NOT read uint16", func(t *testing.T) {
		b.Next(1)

		var val uint16
		err := readUint16(b, binary.LittleEndian, &val)
		if err == nil {
			t.Error("An error should have occurred reading uint16")
		}
	})
}

func TestReadUint32(t *testing.T) {
	d := []byte{0xaa, 0xbb, 0xcc, 0xdd}
	b := bytes.NewBuffer(d)

	t.Run("should read uint32", func(t *testing.T) {
		var val uint32
		err := readUint32(b, binary.LittleEndian, &val)
		if err != nil {
			t.Error("Could not read uint32")
		}

		if val != 0xddccbbaa {
			t.Errorf("Wrong uint32 read: actual (0x%x), expected (0xddccbbaa)", val)
		}
	})

	t.Run("should NOT read uint32", func(t *testing.T) {
		b.Next(1)

		var val uint32
		err := readUint32(b, binary.LittleEndian, &val)
		if err == nil {
			t.Error("An error should have occurred reading uint32")
		}
	})
}

func TestReadUint64(t *testing.T) {
	d := []byte{
		0xaa, 0xbb, 0xcc, 0xdd,
		0x11, 0x22, 0x33, 0x44,
	}
	b := bytes.NewBuffer(d)

	t.Run("should read uint64", func(t *testing.T) {
		var val uint64
		err := readUint64(b, binary.LittleEndian, &val)
		if err != nil {
			t.Error("Could not read uint64")
		}

		if val != 0x44332211ddccbbaa {
			t.Errorf("Wrong uint64 read: actual (0x%x), expected (0x44332211ddccbbaa)", val)
		}
	})

	t.Run("should NOT read uint64", func(t *testing.T) {
		b.Next(1)

		var val uint64
		err := readUint64(b, binary.LittleEndian, &val)
		if err == nil {
			t.Error("An error should have occurred reading uint64")
		}
	})
}

func TestReadSlice(t *testing.T) {
	d := []byte{
		0xaa, 0xbb, 0xcc, 0xdd,
		0x11, 0x22, 0x33, 0x44,
	}

	b := bytes.NewBuffer(d)

	t.Run("should read slice", func(t *testing.T) {
		val := make([]byte, 4)

		err := readSlice(b, binary.BigEndian, &val)
		if err != nil {
			t.Error("Could not read slice")
		}

		expected := d[0:4]
		if bytes.Compare(val, expected) != 0 {
			t.Errorf("Wrong slice read: actual (%v), expected (%v)", val, expected)
		}
	})

	t.Run("should NOT read slice", func(t *testing.T) {
		val := make([]byte, 12)

		err := readSlice(b, binary.BigEndian, &val)
		if err == nil {
			t.Error("An error should have occurred reading slice")
		}
	})
}

func TestDecodeBatch(t *testing.T) {
	d := []byte{
		0xaa, 0xbb, 0xcc, 0xdd,
		0x11, 0x22, 0x33, 0x44,
		0xaf, 0xab, 0xff, 0x6e,
	}

	b := bytes.NewBuffer(d)

	t.Run("should decode batch", func(t *testing.T) {
		byteSliceVal := make([]byte, 4)
		expectedByteSlice := []byte{0xaa, 0xbb, 0xcc, 0xdd}

		var uint16val uint16
		expectUint16Val := uint16(0x1122)

		uint8SliceVal := make([]uint8, 2)
		expectedUInt8Slice := []uint8{0x33, 0x44}

		var uint32val uint32
		expectedUint32val := uint32(0xafabff6e)

		vals := []DecodeVal{
			{
				Order: binary.LittleEndian,
				Val:   &byteSliceVal,
			},
			{
				Order: binary.BigEndian,
				Val:   &uint16val,
			},
			{
				Order: binary.LittleEndian,
				Val:   &uint8SliceVal,
			},
			{
				Order: binary.BigEndian,
				Val:   &uint32val,
			},
		}

		err := DecodeBatch(b, vals...)
		if err != nil {
			t.Error("Could not decode")
		}

		if bytes.Compare(byteSliceVal, expectedByteSlice) != 0 {
			t.Errorf("Wrong slice decoding: actual (%v), expected (%v)", byteSliceVal, expectedByteSlice)
		}

		if uint16val != expectUint16Val {
			t.Errorf("Wrong uint16 decoding: actual (0x%x), expected (0x%x)", uint16val, expectUint16Val)
		}

		if bytes.Compare(uint8SliceVal, expectedUInt8Slice) != 0 {
			t.Errorf("Wrong slice decoding: actual (%v), expected (%v)", uint8SliceVal, expectedUInt8Slice)
		}

		if uint32val != expectedUint32val {
			t.Errorf("Wrong uint32 decoding: actual (0x%x), expected (0x%x)", uint32val, expectedUint32val)
		}
	})
}

func TestWriteUint8(t *testing.T) {
	d := []byte{}
	b := bytes.NewBuffer(d)

	t.Run("should write uint8", func(t *testing.T) {
		var val uint8 = 0xaa
		err := writeUint8(b, binary.LittleEndian, &val)
		if err != nil {
			t.Error("Could not write uint8")
		}

		if b.Bytes()[0] != val {
			t.Errorf("Wrong uint8 write: actual (0x%x), expected (0xaa)", val)
		}
	})
}

func TestWriteUint16(t *testing.T) {
	t.Run("should write uint16 (LittleEndian)", func(t *testing.T) {
		d := []byte{}
		b := bytes.NewBuffer(d)
		var val uint16 = 0xaabb

		err := writeUint16(b, binary.LittleEndian, &val)
		if err != nil {
			t.Error("Could not write uint16")
		}

		expected := [2]byte{
			uint8(val),
			uint8(val >> 8),
		}

		for i, b := range b.Bytes() {
			if expected[i] != b {
				t.Errorf("Wrong b[%d] write: actual (0x%x), expected (0x%x)", i, b, expected[i])
			}
		}
	})

	t.Run("should write uint16 (BigEndian)", func(t *testing.T) {
		d := []byte{}
		b := bytes.NewBuffer(d)
		var val uint16 = 0xaabb

		err := writeUint16(b, binary.BigEndian, &val)
		if err != nil {
			t.Error("Could not write uint16")
		}

		expected := [2]byte{
			uint8(val >> 8),
			uint8(val),
		}

		for i, b := range b.Bytes() {
			if expected[i] != b {
				t.Errorf("Wrong b[%d] write: actual (0x%x), expected (0x%x)", i, b, expected[i])
			}
		}
	})
}

func TestWriteUint32(t *testing.T) {
	t.Run("should write uint32 (LittleEndian)", func(t *testing.T) {
		d := []byte{}
		b := bytes.NewBuffer(d)
		var val uint32 = 0xaabbccdd

		err := writeUint32(b, binary.LittleEndian, &val)
		if err != nil {
			t.Error("Could not write uint32")
		}

		expected := [4]byte{
			uint8(val),
			uint8(val >> 8),
			uint8(val >> 16),
			uint8(val >> 24),
		}

		for i, b := range b.Bytes() {
			if expected[i] != b {
				t.Errorf("Wrong b[%d] write: actual (0x%x), expected (0x%x)", i, b, expected[i])
			}
		}
	})

	t.Run("should write uint32 (BigEndian)", func(t *testing.T) {
		d := []byte{}
		b := bytes.NewBuffer(d)
		var val uint32 = 0xaabbccdd

		err := writeUint32(b, binary.BigEndian, &val)
		if err != nil {
			t.Error("Could not write uint32")
		}

		expected := [4]byte{
			uint8(val >> 24),
			uint8(val >> 16),
			uint8(val >> 8),
			uint8(val),
		}

		for i, b := range b.Bytes() {
			if expected[i] != b {
				t.Errorf("Wrong b[%d] write: actual (0x%x), expected (0x%x)", i, b, expected[i])
			}
		}
	})
}

func TestWriteUint64(t *testing.T) {
	t.Run("should write uint64 (LittleEndian)", func(t *testing.T) {
		d := []byte{}
		b := bytes.NewBuffer(d)
		var val uint64 = 0xaabbccdd11223344

		err := writeUint64(b, binary.LittleEndian, &val)
		if err != nil {
			t.Error("Could not write uint64")
		}

		expected := [8]byte{
			uint8(val),
			uint8(val >> 8),
			uint8(val >> 16),
			uint8(val >> 24),
			uint8(val >> 32),
			uint8(val >> 40),
			uint8(val >> 48),
			uint8(val >> 56),
		}

		for i, b := range b.Bytes() {
			if expected[i] != b {
				t.Errorf("Wrong b[%d] write: actual (0x%x), expected (0x%x)", i, b, expected[i])
			}
		}
	})

	t.Run("should write uint64 (BigEndian)", func(t *testing.T) {
		d := []byte{}
		b := bytes.NewBuffer(d)
		var val uint64 = 0xaabbccdd11223344

		err := writeUint64(b, binary.BigEndian, &val)
		if err != nil {
			t.Error("Could not write uint64")
		}

		expected := [8]byte{
			uint8(val >> 56),
			uint8(val >> 48),
			uint8(val >> 40),
			uint8(val >> 32),
			uint8(val >> 24),
			uint8(val >> 16),
			uint8(val >> 8),
			uint8(val),
		}

		for i, b := range b.Bytes() {
			if expected[i] != b {
				t.Errorf("Wrong b[%d] write: actual (0x%x), expected (0x%x)", i, b, expected[i])
			}
		}
	})
}

func TestEncodeBatch(t *testing.T) {
	d := []byte{}
	b := bytes.NewBuffer(d)

	t.Run("should encode batch", func(t *testing.T) {
		var uint32Val uint32 = 0xafabff6e
		var uint64Val uint64 = 0x00000506ffbb
		slice := []byte{0xaa, 0x8b, 0x9c}

		vals := []EncodeVal{
			{
				Order: binary.LittleEndian,
				Val:   &uint32Val,
			},
			{
				Order: binary.BigEndian,
				Val:   &uint64Val,
			},
			{
				Order: binary.LittleEndian,
				Val:   &slice,
			},
		}

		err := EncodeBatch(b, vals...)
		if err != nil {
			t.Error("Could not encode")
		}

		expected := []byte{
			uint8(uint32Val),
			uint8(uint32Val >> 8),
			uint8(uint32Val >> 16),
			uint8(uint32Val >> 24),
			uint8(uint64Val >> 56),
			uint8(uint64Val >> 48),
			uint8(uint64Val >> 40),
			uint8(uint64Val >> 32),
			uint8(uint64Val >> 24),
			uint8(uint64Val >> 16),
			uint8(uint64Val >> 8),
			uint8(uint64Val),
		}

		expected = append(expected, slice...)

		if bytes.Compare(b.Bytes(), expected) != 0 {
			t.Error("Wrong encoding")
		}
	})
}
