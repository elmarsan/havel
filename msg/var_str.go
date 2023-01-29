package msg

import (
	"encoding/binary"
	"io"
)

// VarStr represents variable length string using in messaging.
// https://en.bitcoin.it/wiki/Protocol_documentation#Variable_length_string
type VarStr struct {
	// VarInt represents the VarInt of the string.
	VarInt VarInt
	// Val holds the string.
	Val string
}

// Decode decodes Str from r.
func (varStr *VarStr) Decode(r io.Reader) error {
	// Decode len
	err := varStr.VarInt.Decode(r)
	if err != nil {
		return err
	}

	// Allocate space for str using Len
	val := make([]byte, varStr.VarInt.Length)
	err = Decode(r, binary.LittleEndian, &val)
	if err != nil {
		return err
	}

	varStr.Val = string(val)
	return nil
}

// Encode encodes MsgNetAddr into w.
func (varStr *VarStr) Encode(w io.Writer) error {
	// Encode length
	err := varStr.VarInt.Encode(w)
	if err != nil {
		return err
	}

	// Encode string
	val := make([]byte, varStr.VarInt.Length)
	copy(val[:], varStr.Val)

	err = Encode(w, binary.LittleEndian, &val)
	if err != nil {
		return err
	}

	return nil
}
