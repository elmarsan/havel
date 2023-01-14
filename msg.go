package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// Msg represents a message used in Bitcoin p2p communication.
type Msg interface {
	Encode(b *bytes.Buffer) error
	Decode(b *bytes.Buffer) error
}

const (
	MsgHeaderSize = 24 // MsgHeaderSize represents maximun header size allowed
	CmdSize       = 12 // CmdSize represents maximun cmd size allowed
)

// https://en.bitcoin.it/wiki/Protocol_documentation#Message_structure
type MsgHeader struct {
	Magic    BitcoinNet // Magic represents the value indicating the origin network.
	Cmd      BitcoinCmd // Cmd represents the content type of the msg.
	Length   uint32     // Lenght of payload in number of bytes.
	Checksum uint32     // Checksum holds first 4 bytes of sha256(sha256(payload)).
}

// // Encode writes MsgHeader properties to bytes Buffer
// // It overwritte the first 24 bytes of the buffer
// func (msgHeader *MsgHeader) Encode(b *bytes.Buffer) error {
// 	return nil
// }

// Decode initializes MsgHeader properties from bytes Buffer
func (msgHeader *MsgHeader) Decode(b *bytes.Buffer) error {
	magicUint32 := binary.LittleEndian.Uint32(b.Bytes()[0:4])
	magic, err := NewBitcoinNet(magicUint32)
	if err != nil {
		return err
	}

	hexCmd := "0x" + hex.EncodeToString(b.Bytes()[4:16])
	cmd, err := NewBitcoinCmd(hexCmd)
	if err != nil {
		return err
	}

	length := binary.LittleEndian.Uint32(b.Bytes()[16:20])
	checksum := binary.BigEndian.Uint32(b.Bytes()[20:24])

	msgHeader.Magic = *magic
	msgHeader.Cmd = *cmd
	msgHeader.Length = length
	msgHeader.Checksum = checksum

	return nil
}
