package msg

import (
	"bytes"
	"net"
	"reflect"
	"testing"
)

func TestNetAddr(t *testing.T) {
	data := []byte{
		// Services
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		// Ip
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0xff, 0xff, 0x5d, 0xb0, 0x82,
		0x8b,
		// Port
		0x93, 0x59,
	}

	sample := &NetAddr{
		Services: 0x00000000,
		Ip: net.IP{
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0xff, 0xff, 0x5d, 0xb0, 0x82, 0x8b,
		},
		Port: 0x9359,
	}

	t.Run("Decode", func(t *testing.T) {
		b := bytes.NewBuffer(data)

		netAddr := &NetAddr{}
		err := netAddr.Decode(b)
		if err != nil {
			t.Error("Could not decode")
		}

		if !reflect.DeepEqual(netAddr, sample) {
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
