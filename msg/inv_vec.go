package msg

import (
	"encoding/binary"
	"fmt"
	"io"
)

type InvObj uint32

const (
	UNKNOWN                    InvObj = 0
	MSG_TX                     InvObj = 1
	MSG_BLOCK                  InvObj = 2
	MSG_FILTERED_BLOCK         InvObj = 3
	MSG_CMPCT_BLOCK            InvObj = 4
	MSG_WITNESS_TX             InvObj = 0x40000001
	MSG_WITNESS_BLOCK          InvObj = 0x40000002
	MSG_FILTERED_WITNESS_BLOCK InvObj = 0x40000003
)

// btcNetUint32 is a map of inventory objects back to their uint32 value
var invObjUint32 = map[uint32]InvObj{
	uint32(UNKNOWN):                    UNKNOWN,
	uint32(MSG_TX):                     MSG_TX,
	uint32(MSG_BLOCK):                  MSG_BLOCK,
	uint32(MSG_FILTERED_BLOCK):         MSG_FILTERED_BLOCK,
	uint32(MSG_CMPCT_BLOCK):            MSG_CMPCT_BLOCK,
	uint32(MSG_WITNESS_TX):             MSG_WITNESS_TX,
	uint32(MSG_WITNESS_BLOCK):          MSG_WITNESS_BLOCK,
	uint32(MSG_FILTERED_WITNESS_BLOCK): MSG_FILTERED_WITNESS_BLOCK,
}

// NewInvObj returns InvObj matching the uint32 value.
func NewInvObj(invObj uint32) (*InvObj, error) {
	if s, ok := invObjUint32[invObj]; ok {
		return &s, nil
	}

	return nil, fmt.Errorf("Unknown InvObj (%d)", invObj)
}

// https://en.bitcoin.it/wiki/Protocol_documentation#Inventory_Vectors
type InvVec struct {
	// Obj represents obj type.
	Obj InvObj
	// Hash represents obj hash.
	Hash [32]byte
}

// Decode decodes InvVec from r.
func (iv *InvVec) Decode(r io.Reader) error {
	hash := make([]byte, 32)
	var obj uint32

	vals := []DecodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &obj,
		},
		{
			Order: binary.LittleEndian,
			Val:   &hash,
		},
	}

	err := DecodeBatch(r, vals...)
	if err != nil {
		return err
	}

	invObj, err := NewInvObj(obj)
	if err != nil {
		return err
	}

	iv.Obj = *invObj

	copy(iv.Hash[:], hash)

	return nil
}

// Encode encodes InvVec into w.
func (iv *InvVec) Encode(w io.Writer) error {
	hash := make([]byte, 32)
	copy(hash[:], iv.Hash[:])

	obj := uint32(iv.Obj)

	vals := []EncodeVal{
		{
			Order: binary.LittleEndian,
			Val:   &obj,
		},
		{
			Order: binary.LittleEndian,
			Val:   &hash,
		},
	}

	err := EncodeBatch(w, vals...)
	if err != nil {
		return err
	}

	return nil
}
