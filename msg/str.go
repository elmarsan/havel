package msg

import (
	"encoding/binary"
	"io"
)

// Str represents variable length string using in messaging.
// https://en.bitcoin.it/wiki/Protocol_documentation#Variable_length_string
type Str struct {
	// Len represents the Len of the string.
	Len uint8
	// Val holds the string.
	Val string
}

// Decode decodes Str from r.
func (str *Str) Decode(r io.Reader) error {
	// Decode len
	err := Decode(r, binary.LittleEndian, &str.Len)
	if err != nil {
		return err
	}

	// Allocate space for str using Len
	val := make([]byte, str.Len)
	err = Decode(r, binary.LittleEndian, &val)
	if err != nil {
		return err
	}

	str.Val = string(val)

	return nil
}

// Encode encodes MsgNetAddr into w.
func (str *Str) Encode(w io.Writer) error {
	val := make([]byte, str.Len)
	copy(val[:], str.Val)

	vals := []EncodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &str.Len,
		},
		{
			Order: binary.LittleEndian,
			Val:   &val,
		},
	}

	return EncodeBatch(w, vals...)
}
