package msg

import (
	"bytes"
	"reflect"
	"testing"
)

func TestInvVec(t *testing.T) {
	data := []byte{
		0x02, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xcb, 0x1f, 0x16, 0xba, 0x92, 0x3a,
		0x07, 0x3b, 0x79, 0xc3, 0xdf, 0x15, 0x8a, 0x92,
		0x40, 0xd5, 0x2c, 0xa8, 0x1b, 0xba, 0x18, 0xbc,
	}

	hash := [32]byte{}
	copy(hash[:], data[4:])

	sample := &InvVec{
		Obj:  MSG_BLOCK,
		Hash: hash,
	}

	t.Run("Decode", func(t *testing.T) {
		b := bytes.NewBuffer(data)

		invVec := &InvVec{}
		err := invVec.Decode(b)
		if err != nil {
			t.Errorf("Unable to decode (%s)", err)
		}

		if !reflect.DeepEqual(invVec, sample) {
			t.Error("Wrong decoding")
		}
	})

	t.Run("Encode", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})

		err := sample.Encode(b)
		if err != nil {
			t.Errorf("Unable to encode (%s)", err)
		}

		if bytes.Compare(b.Bytes(), data) != 0 {
			t.Error("Wrong encoding")
		}
	})
}
