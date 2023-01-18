package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

// MsgData represents message data used in Bitcoin p2p communication.
type MsgData interface {
	Encode(b *bytes.Buffer) error
	Decode(b *bytes.Buffer) error
}

const (
	MsgHeaderSize = 24 // MsgHeaderSize represents maximun header size allowed
	CmdSize       = 12 // CmdSize represents maximun cmd size allowed
)

// MsgHeader represents a default information contained in every protocol msg.
// https://en.bitcoin.it/wiki/Protocol_documentation#Message_structure
type MsgHeader struct {
	Magic    BitcoinNet // Magic represents the value indicating the origin network.
	Cmd      BitcoinCmd // Cmd represents the content type of the msg.
	Length   uint32     // Lenght of payload in number of bytes.
	Checksum uint32     // Checksum holds first 4 bytes of sha256(sha256(payload)).
}

// Encode encodes MsgHeader in given Buffer.
// TODO: Implement
func (msgHeader *MsgHeader) Encode(b *bytes.Buffer) error {
	return nil
}

// Decode decodes MsgHeader contained in given buffer.
// [0:4] BitcoinNet
// [4:16] BitcoinCmd
// [16:20] Checksum
// [20:24] Length
func (msgHeader *MsgHeader) Decode(b *bytes.Buffer) error {
	// Decode magic
	magicBytes := make([]byte, 4)
	_, err := b.Read(magicBytes)
	if err != nil {
		return fmt.Errorf("Could not decode magic")
	}

	magicUint32 := binary.LittleEndian.Uint32(magicBytes)
	magic, err := NewBitcoinNet(magicUint32)
	if err != nil {
		return err
	}

	msgHeader.Magic = *magic

	// Decode command
	cmdBytes := make([]byte, 12)
	_, err = b.Read(cmdBytes)
	if err != nil {
		return fmt.Errorf("Could not decode command")
	}

	hexCmd := "0x" + hex.EncodeToString(cmdBytes)
	cmd, err := NewBitcoinCmd(hexCmd)
	if err != nil {
		return err
	}

	msgHeader.Cmd = *cmd

	// Decode length
	lenghtBytes := make([]byte, 4)
	_, err = b.Read(lenghtBytes)
	if err != nil {
		return fmt.Errorf("Could not decode length")
	}
	length := binary.LittleEndian.Uint32(lenghtBytes)
	msgHeader.Length = length

	// Decode checksum
	checksumBytes := make([]byte, 4)
	_, err = b.Read(checksumBytes)
	if err != nil {
		return fmt.Errorf("Could not decode checksum")
	}
	checksum := binary.BigEndian.Uint32(checksumBytes)
	msgHeader.Checksum = checksum

	return nil
}

// MsgNetAddr represents network address of node.
// https://en.bitcoin.it/wiki/Protocol_documentation#Network_address
type MsgNetAddr struct {
	// Timestamp represents standard UNIX timestamp.
	Timestamp time.Time
	// Services represents bitfield of features to be enabled for this connection.
	Services uint64
	// Ip represents node's ip.
	Ip net.IP
	// Port represent node's port.
	Port uint16
}

// Decode decodes MsgNetAddr contained in given buffer.
// [0:4] Timestamp
// [4:12] Services
// [12:28] IP
// [28:30] Port
func (msgNetAddr *MsgNetAddr) Decode(b *bytes.Buffer) error {
	// TODO: Check size taking into account msg version does not include timestamp
	unixBytes := make([]byte, 4)
	_, err := b.Read(unixBytes)
	if err != nil {
		return fmt.Errorf("Could not decode timestamp")
	}

	unix := binary.LittleEndian.Uint32(unixBytes)
	msgNetAddr.Timestamp = time.Unix(int64(unix), 0)

	// TODO: Add service
	b.Next(8)

	ipBytes := make([]byte, 16)
	_, err = b.Read(ipBytes)
	if err != nil {
		return fmt.Errorf("Could not decode ip")
	}

	msgNetAddr.Ip = net.IP(ipBytes)

	portBytes := make([]byte, 2)
	_, err = b.Read(portBytes)
	if err != nil {
		return fmt.Errorf("Could not decode port")
	}

	msgNetAddr.Port = binary.BigEndian.Uint16(portBytes)

	return nil
}

// Encode encodes MsgNetAddr in given buffer.
// [0:4] BitcoinNet
// [4:16] BitcoinCmd
// [16:20] Checksum
// [20:24] Length
func (msgNetAddr *MsgNetAddr) Encode(b *bytes.Buffer) error {

	return nil
}
