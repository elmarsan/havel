package main

import (
	"bytes"
	"log"
	"net"
	"reflect"
	"testing"
)

func TestMsgHeader(t *testing.T) {
	versionMsgHeader := []byte{
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

	t.Run("Decode", func(t *testing.T) {
		buf := bytes.NewBuffer(versionMsgHeader)

		msgHeader := &MsgHeader{}
		msgHeader.Decode(buf)

		log.Println(msgHeader.Length)

		if *(msgHeader).Magic != MainNet {
			t.Errorf("Magic (0x%x), expected (0x%x)", msgHeader.Magic, MainNet)
		}

		if msgHeader.Cmd.Name != VersionCmd {
			t.Errorf("Cmd (0x%x), expected (%s)", msgHeader.Cmd, VersionCmd)
		}

		length := 0x64
		if msgHeader.Length != uint32(length) {
			t.Errorf("Lenght (0x%x), expected (0x%x)", msgHeader.Length, length)
		}

		cheksum := 0x358d4932
		if msgHeader.Checksum != uint32(cheksum) {
			t.Errorf("Checksum (0x%x), expected (0x%x)", msgHeader.Checksum, cheksum)
		}
	})

	t.Run("Encode", func(t *testing.T) {
		mainnet := MainNet

		msgHeader := &MsgHeader{
			Magic: &mainnet,
			Cmd: &BitcoinCmd{
				HexData: VersionCmdData,
				Name:    VersionCmd,
			},
			Length:   0x64000000,
			Checksum: 0x358d4932,
		}

		d := []byte{}
		b := bytes.NewBuffer(d)

		err := msgHeader.Encode(b)
		if err != nil {
			t.Errorf("Wrong msg header encoding, cause (%s)", err.Error())
		}

		versionMsgHeader := []byte{
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

		if bytes.Compare(b.Bytes(), versionMsgHeader) != 0 {
			t.Error("Wrong msg header encoding")
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

	t.Run("Decode", func(t *testing.T) {
		b := bytes.NewBuffer(data)

		expected := &MsgNetAddr{
			Services: 0x00000000,
			Ip: net.IP{
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0xff, 0xff, 0x5d, 0xb0, 0x82, 0x8b,
			},
			Port: 0x9359,
		}

		msgNetAddr := &MsgNetAddr{}
		err := msgNetAddr.Decode(b)
		if err != nil {
			t.Error("Could not decode")
		}

		if !reflect.DeepEqual(msgNetAddr, expected) {
			t.Error("Wrong decoding")
		}
	})

	t.Run("Encode", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})

		msgNetAddr := &MsgNetAddr{
			Services: 0x00000000,
			Ip: net.IP{
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0xff, 0xff, 0x5d, 0xb0, 0x82, 0x8b,
			},
			Port: 0x9359,
		}

		err := msgNetAddr.Encode(b)
		if err != nil {
			t.Error("Could not encode")
		}

		if bytes.Compare(b.Bytes(), data) != 0 {
			t.Errorf("Wrong encoding")
		}
	})
}
