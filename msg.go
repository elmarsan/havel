package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
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

// MsgNetAddr represents network address of node.
// https://en.bitcoin.it/wiki/Protocol_documentation#Network_address
type MsgNetAddr struct {
	// Timestamp represents standard UNIX timestamp.
	Timestamp time.Time
	// Services represents bitfield of features to be enabled for this connection.
	Services uint64
	// Addr represents the address of node.
	Addr net.Addr
}

// Decode decodes MsgNetAddr contained in given buffer.
// [0:4] Timestamp
// [4:12] Services
// [12:28] IP
// [28:30] Port
func (msgNetAddr *MsgNetAddr) Decode(b *bytes.Buffer) error {
	// TODO: Check size taking into account msg version does not include timestamp
	unix := binary.LittleEndian.Uint32(b.Bytes()[0:4])

	fmt.Println(unix)

	msgNetAddr.Timestamp = time.Unix(int64(unix), 0)

	// TODO: Add service
	// services := binary.LittleEndian.Uint64(b.Bytes()[4:12])
	// msgNetAddr.Services = services

	ip := net.IP(b.Bytes()[12:28])
	port := binary.BigEndian.Uint16(b.Bytes()[28:30])

	tcpAddr := net.JoinHostPort(ip.String(), strconv.Itoa(int(port)))
	addr, err := net.ResolveTCPAddr("tcp", tcpAddr)
	if err != nil {
		return err
	}

	msgNetAddr.Addr = addr

	return nil
}

// Decode initializes MsgNetAddr properties from bytes Buffer
// [0:4] BitcoinNet
// [4:16] BitcoinCmd
// [16:20] Checksum
// [20:24] Length
func (msgNetAddr *MsgNetAddr) Encode(b *bytes.Buffer) error {
	data := []byte{}
	unix := uint32(msgNetAddr.Timestamp.Unix())
	binary.LittleEndian.PutUint32(data, unix)

	services := []byte{}
	binary.LittleEndian.PutUint64(services, msgNetAddr.Services)
	data = append(data, services...)

	// TODO: add ip and port

	_, err := b.Write(data)
	if err != nil {
		return err
	}

	// TODO: Check n and compare with expected size
	return nil
}
