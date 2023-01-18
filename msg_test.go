package main

import (
	"bytes"
	"testing"
	"time"
)

func TestMsgHeaderDecode(t *testing.T) {
	t.Run("should decode and initialize properties", func(t *testing.T) {
		versionMsgHeader := []byte{
			0xf9, 0xbe, 0xb4, 0xd9, 0x76, 0x65, 0x72, 0x73,
			0x69, 0x6f, 0x6e, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x64, 0x00, 0x00, 0x00, 0x35, 0x8d, 0x49, 0x32,
		}

		buf := bytes.NewBuffer(versionMsgHeader)
		msgHeader := &MsgHeader{}
		msgHeader.Decode(buf)

		if msgHeader.Magic != MainNet {
			t.Errorf("Magic (0x%x), expected (0x%x)", msgHeader.Magic, MainNet)
		}

		if msgHeader.Cmd != VersionCmd {
			t.Errorf("Cmd (0x%x), expected (%s)", msgHeader.Cmd, VersionCmdHex)
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
}

func TestMsgNetAddr(t *testing.T) {
	t.Run("Decode", func(t *testing.T) {
		data := []byte{
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
			0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff,
			0x5d, 0xb0, 0x82, 0x8b, 0x93, 0x59,
		}

		b := bytes.NewBuffer(data)

		msgNetAddr := &MsgNetAddr{}
		msgNetAddr.Decode(b)

		ip := msgNetAddr.Ip.String()
		expectedIp := "93.176.130.139"
		if ip != expectedIp {
			t.Errorf("Wrong ip (%s), expected (%s)", ip, expectedIp)
		}

		port := int(msgNetAddr.Port)
		expectedPort := 37721
		if port != expectedPort {
			t.Errorf("Wrong port (%d), expected (%d)", port, expectedPort)
		}

		if msgNetAddr.Timestamp != time.Unix(int64(0), 0) {
			t.Errorf("Wrong timestamp")
		}
	})
}
