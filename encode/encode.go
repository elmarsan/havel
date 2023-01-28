package encode

import (
	"encoding/binary"
	"io"
)

// Encoder defines signature for encoding and decoding data.
type Encoder interface {
	// Encode encodes data into w
	Encode(w io.Writer) error
	// Decode decode data from r
	Decode(r io.Reader) error
}

// readUint8 reads uint8 from r into d.
func readUint8(r io.Reader, order binary.ByteOrder, d *uint8) error {
	if err := binary.Read(r, order, d); err != nil {
		return err
	}

	return nil
}

// readUint16 reads uint16 from r into d.
func readUint16(r io.Reader, order binary.ByteOrder, d *uint16) error {
	if err := binary.Read(r, order, d); err != nil {
		return err
	}

	return nil
}

// readUint32 reads uint32 from r into d.
func readUint32(r io.Reader, order binary.ByteOrder, d *uint32) error {
	if err := binary.Read(r, order, d); err != nil {
		return err
	}

	return nil
}

// readUint64 reads uint64 from r into d.
func readUint64(r io.Reader, order binary.ByteOrder, d *uint64) error {
	if err := binary.Read(r, order, d); err != nil {
		return err
	}

	return nil
}

// readSlice reads slice from r into d.
func readSlice(r io.Reader, order binary.ByteOrder, d *[]byte) error {
	if err := binary.Read(r, order, d); err != nil {
		return err
	}

	return nil
}

// Decode decodes r into d.
func Decode(r io.Reader, order binary.ByteOrder, d interface{}) error {
	switch t := d.(type) {
	case *uint8:
		return readUint8(r, order, t)
	case *uint16:
		return readUint16(r, order, t)
	case *uint32:
		return readUint32(r, order, t)
	case *uint64:
		return readUint64(r, order, t)
	case *[]byte:
		return readSlice(r, order, t)
	}
	return nil
}

// DecodeVal holds target value for decoding.
type DecodeVal struct {
	// Order specifies encoding byte Order.
	Order binary.ByteOrder
	// Val holds data of val.
	Val interface{}
}

// DecodeBatch decodes batch of values from b into DecodeVal.Val.
func DecodeBatch(r io.Reader, vals ...DecodeVal) error {
	for _, val := range vals {
		if err := Decode(r, val.Order, val.Val); err != nil {
			return err
		}
	}

	return nil
}

// writeUint8 writes uint8 from w into d.
func writeUint8(w io.Writer, order binary.ByteOrder, d *uint8) error {
	if err := binary.Write(w, order, d); err != nil {
		return err
	}

	return nil
}

// writeUint16 writes uint16 from w into d.
func writeUint16(w io.Writer, order binary.ByteOrder, d *uint16) error {
	if err := binary.Write(w, order, d); err != nil {
		return err
	}

	return nil
}

// writeUint32 writes uint32 from w into d.
func writeUint32(w io.Writer, order binary.ByteOrder, d *uint32) error {
	if err := binary.Write(w, order, d); err != nil {
		return err
	}

	return nil
}

// writeUint64 writes uint64 from w into d.
func writeUint64(w io.Writer, order binary.ByteOrder, d *uint64) error {
	if err := binary.Write(w, order, d); err != nil {
		return err
	}

	return nil
}

// writeSlice writes slice from w into d.
func writeSlice(w io.Writer, order binary.ByteOrder, d *[]byte) error {
	if err := binary.Write(w, order, d); err != nil {
		return err
	}

	return nil
}

// encode encodes d into w.
func encode(w io.Writer, order binary.ByteOrder, d interface{}) error {
	switch t := d.(type) {
	case *uint8:
		return writeUint8(w, order, t)
	case *uint16:
		return writeUint16(w, order, t)
	case *uint32:
		return writeUint32(w, order, t)
	case *uint64:
		return writeUint64(w, order, t)
	case *[]byte:
		return writeSlice(w, order, t)
	}

	return nil
}

// EncodeVal holds target value for encoding.
type EncodeVal struct {
	// Order specifies decoding byte Order.
	Order binary.ByteOrder
	// Val holds data of val.
	Val interface{}
}

// EncodeBatch encodes batch of values from EncodeVal.Val into w.
func EncodeBatch(w io.Writer, vals ...EncodeVal) error {
	for _, val := range vals {
		if err := encode(w, val.Order, val.Val); err != nil {
			return err
		}
	}

	return nil
}
