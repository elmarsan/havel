package main

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/elmarsan/havel/msg"
	"github.com/elmarsan/havel/protocol"
)

// Client represents Bitcoin network client
type Client struct {
	// version represents the protocol version used by the node.
	version uint32
	// net represents Bitcoin network (mainnet, testnet, etc...)
	net protocol.BitcoinNet

	// peers represents client connected peers.
	peers []Peer
}

// Peer represents Bitcoin network node.
type Peer struct {
	// conn holds the connection to the peer.
	conn net.Conn
}

// TODO: Implement
// 1 - Connected to peer
// 2 - Send version msg
// 3 - Send verack msg
// 4 - Save peer and keep dial openned
func (c *Client) AddPeer(addr string) error {
	fields := strings.Split(addr, ":")
	ip := net.ParseIP(addr)
	port, err := strconv.Atoi(fields[1])
	if err != nil {
		return err
	}

	version := msg.Version{
		Header: &msg.Header{
			Magic: c.net,
			Cmd: protocol.BitcoinCmd{
				HexData: protocol.VersionCmdData,
				Name:    protocol.VersionCmd,
			},
			Length:   0x64,
			Checksum: 0x358d4932,
		},
		Version:   c.version,
		Services:  0x00000001,
		Timestamp: time.Now(),
		Nonce:     0x6517e68c5db32e3b,
		RecvAddr: &msg.NetAddr{
			Ip:       ip,
			Port:     uint16(port),
			Services: 1,
		},
		FromAddr: &msg.NetAddr{
			Ip:       ip,
			Port:     uint16(port),
			Services: 1,
		},
		UserAgent: &msg.VarStr{
			VarInt: msg.VarInt{
				Length: 13,
			},
			Val: "/Havel:0.0.1/",
		},
		StartHeight: 0,
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("Unable to connect peer (%s)", err)
	}

	defer conn.Close()

	versionData := bytes.NewBuffer([]byte{})
	err = version.Encode(versionData)
	if err != nil {
		return err
	}

	_, err = conn.Write(versionData.Bytes())
	if err != nil {
		return fmt.Errorf("Unable to send version: %v\n", err)
	}

	// _, err := ioutil.ReadAll(conn)
	// if err != nil {
	// 	return err
	// }

	return nil
}
