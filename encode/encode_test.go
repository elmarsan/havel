package encode

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
			t.Error("Could not read uint32")
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
				order: binary.LittleEndian,
				d:     &byteSliceVal,
			},
			{
				order: binary.BigEndian,
				d:     &uint16val,
			},
			{
				order: binary.LittleEndian,
				d:     &uint8SliceVal,
			},
			{
				order: binary.BigEndian,
				d:     &uint32val,
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
