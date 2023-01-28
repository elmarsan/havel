package main

import (
	"bytes"
	"net"
	"reflect"
	"testing"
)

func TestMsgHeader(t *testing.T) {
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

	mainnet := MainNet

	sample := &MsgHeader{
		Magic: &mainnet,
		Cmd: &BitcoinCmd{
			HexData: VersionCmdData,
			Name:    VersionCmd,
		},
		Length:   0x64,
		Checksum: 0x32498d35,
	}

	t.Run("Decode", func(t *testing.T) {
		b := bytes.NewBuffer(data)

		header := &MsgHeader{}
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

func TestMsgNetAddr(t *testing.T) {
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

	sample := &MsgNetAddr{
		Services: 0x00000000,
		Ip: net.IP{
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0xff, 0xff, 0x5d, 0xb0, 0x82, 0x8b,
		},
		Port: 0x9359,
	}

	t.Run("Decode", func(t *testing.T) {
		b := bytes.NewBuffer(data)

		msgNetAddr := &MsgNetAddr{}
		err := msgNetAddr.Decode(b)
		if err != nil {
			t.Error("Could not decode")
		}

		if !reflect.DeepEqual(msgNetAddr, sample) {
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

func TestMsgStr(t *testing.T) {
	data := []byte{
		// "/Satoshi:0.7.2/"
		0x0F, 0x2F, 0x53, 0x61, 0x74, 0x6F, 0x73, 0x68,
		0x69, 0x3A, 0x30, 0x2E, 0x37, 0x2E, 0x32, 0x2F,
	}

	sample := &MsgStr{
		Len: 15,
		Str: "/Satoshi:0.7.2/",
	}

	t.Run("Decode", func(t *testing.T) {
		b := bytes.NewBuffer(data)

		msgStr := &MsgStr{}
		err := msgStr.Decode(b)
		if err != nil {
			t.Error("Could not decode")
		}

		if !reflect.DeepEqual(msgStr, sample) {
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
