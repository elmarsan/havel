package msg

import (
	"bytes"
	"reflect"
	"testing"
)

func TestStr(t *testing.T) {
	data := []byte{
		// "/Satoshi:0.7.2/"
		0x0F, 0x2F, 0x53, 0x61, 0x74, 0x6F, 0x73, 0x68,
		0x69, 0x3A, 0x30, 0x2E, 0x37, 0x2E, 0x32, 0x2F,
	}

	sample := &VarStr{
		VarInt: VarInt{
			Length: 15,
		},
		Val: "/Satoshi:0.7.2/",
	}

	t.Run("Decode", func(t *testing.T) {
		b := bytes.NewBuffer(data)

		str := &VarStr{}
		err := str.Decode(b)
		if err != nil {
			t.Error("Could not decode")
		}

		if !reflect.DeepEqual(str, sample) {
			t.Error("Wrong decoding")
		}
	})

	t.Run("Encode", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})

		err := sample.Encode(b)
		if err != nil {
			t.Error("Could not encode")
		}

		if bytes.Compare(b.Bytes(), data) != 0 {
			t.Errorf("Wrong encoding")
		}
	})
}
