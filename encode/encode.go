package encode

import (
	"bytes"
	"encoding/binary"
)

// readUint8 reads uint8 from b into d.
func readUint8(b *bytes.Buffer, order binary.ByteOrder, d *uint8) error {
	if err := binary.Read(b, order, d); err != nil {
		return err
	}

	return nil
}

// readUint16 reads uint16 from b into d.
func readUint16(b *bytes.Buffer, order binary.ByteOrder, d *uint16) error {
	if err := binary.Read(b, order, d); err != nil {
		return err
	}

	return nil
}

// readUint32 reads uint32 from b into d.
func readUint32(b *bytes.Buffer, order binary.ByteOrder, d *uint32) error {
	if err := binary.Read(b, order, d); err != nil {
		return err
	}

	return nil
}

// readUint64 reads uint64 from b into d.
func readUint64(b *bytes.Buffer, order binary.ByteOrder, d *uint64) error {
	if err := binary.Read(b, order, d); err != nil {
		return err
	}

	return nil
}

// readSlice reads uint64 from b into d.
func readSlice(b *bytes.Buffer, order binary.ByteOrder, d *[]byte) error {
	if err := binary.Read(b, order, d); err != nil {
		return err
	}

	return nil
}

// decode decodes any type from b into d.
func decode(b *bytes.Buffer, order binary.ByteOrder, d interface{}) error {
	switch t := d.(type) {
	case *uint8:
		return readUint8(b, order, t)
	case *uint16:
		return readUint16(b, order, t)
	case *uint32:
		return readUint32(b, order, t)
	case *uint64:
		return readUint64(b, order, t)
	case *[]byte:
		return readSlice(b, order, t)
	}
	return nil
}

// DecodeVal holds target value for decoding.
type DecodeVal struct {
	// order specifies encoding byte order
	order binary.ByteOrder
	// d holds data of val
	d interface{}
}

// DecodeBatch decodes batch of values from b into DecodeVal.d
func DecodeBatch(b *bytes.Buffer, vals ...DecodeVal) error {
	for _, val := range vals {
		if err := decode(b, val.order, val.d); err != nil {
			return err
		}
	}

	return nil
}
