package main

import (
	"fmt"
	"strings"
)

// BitcoinNet represents which bitcoin network a message belongs to.
type BitcoinNet uint32

// Constants used to indicate the message bitcoin network.
const (
	MainNet  BitcoinNet = 0xd9b4bef9 // MainNet represents the main bitcoin network.
	TestNet  BitcoinNet = 0xdab5bffa // TestNet represents the regression test network.
	TestNet3 BitcoinNet = 0x0709110b // TestNet3 represents the test network (version 3).
	SimNet   BitcoinNet = 0x12141c16 // SimNet represents the simulation test network.
)

// btcNetUint32 is a map of bitcoin networks back to their uint32 value
var btcNetUint32 = map[uint32]BitcoinNet{
	uint32(MainNet):  MainNet,
	uint32(TestNet):  TestNet,
	uint32(TestNet3): TestNet3,
	uint32(SimNet):   SimNet,
}

// NewBitcoinNet returns BitcoinNet matching the uint32 value.
func NewBitcoinNet(net uint32) (*BitcoinNet, error) {
	if s, ok := btcNetUint32[net]; ok {
		return &s, nil
	}

	return nil, fmt.Errorf("Unknown BitcoinNet (%d)", net)
}

// BitcoinCmd represents a command used in p2p communication.
type BitcoinCmd [12]byte

// VersionCmd represents version command of bitcoin protocol.
var VersionCmd BitcoinCmd = BitcoinCmd{0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x00, 0x00, 0x00, 0x00, 0x00}

// VersionCmdHex represents version command of bitcoin protocol in hexadecimal string value.
var VersionCmdHex string = "0x76657273696f6e0000000000"

// btcNetUint32 is a map of bitcoin networks back to their uint32 value
var btcCmdHexString = map[string]BitcoinCmd{
	VersionCmdHex: VersionCmd,
}

// NewBitcoinCmd returns BitcoinCmd matching the hexadecimal string value.
func NewBitcoinCmd(cmdHex string) (*BitcoinCmd, error) {
	if s, ok := btcCmdHexString[strings.ToLower(cmdHex)]; ok {
		return &s, nil
	}

	return nil, fmt.Errorf("Unknown BitcoinCmd (%s)", cmdHex)
}
