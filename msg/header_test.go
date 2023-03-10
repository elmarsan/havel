package msg

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/elmarsan/havel/protocol"
)

func TestHeader(t *testing.T) {
	data := []byte{
		// Magic
		0xf9, 0xbe, 0xb4, 0xd9,
		// Command
		0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x00,
		0x00, 0x00, 0x00, 0x00,
		// Length
		0x64, 0x00, 0x00, 0x00,
		// Checksum
		0x35, 0x8d, 0x49, 0x32,
	}

	sample := &Header{
		Magic: protocol.MainNet,
		Cmd: protocol.BitcoinCmd{
			HexData: protocol.VersionCmdData,
			Name:    protocol.VersionCmd,
		},
		Length:   0x64,
		Checksum: 0x32498d35,
	}

	t.Run("Decode", func(t *testing.T) {
		b := bytes.NewBuffer(data)

		header := &Header{}
		err := header.Decode(b)
		if err != nil {
			t.Errorf("Unable to decode (%s)", err.Error())
		}

		if !reflect.DeepEqual(header, sample) {
			t.Error("Wrong encoding")
		}
	})

	t.Run("Encode", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})

		err := sample.Encode(b)
		if err != nil {
			t.Errorf("Unable to encode (%s)", err.Error())
		}

		if bytes.Compare(b.Bytes(), data) != 0 {
			t.Error("Wrong encoding")
		}
	})
}
