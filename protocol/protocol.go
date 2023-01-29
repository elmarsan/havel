package protocol

import (
	"fmt"
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

// BitcoinCmdSize represents the size of p2p cmd.
const BitcoinCmdSize = 12

// BitcoinCmdData represents command data shared between nodes.
type BitcoinCmdData [BitcoinCmdSize]byte

// BitcoinCmdName represents a command used in p2p communication.
type BitcoinCmdName string

const (
	VersionCmd BitcoinCmdName = "Version"
)

var VersionCmdData BitcoinCmdData = BitcoinCmdData{0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x00, 0x00, 0x00, 0x00, 0x00}

// btcCmdDataName is a map of BitcoinCmdData back to their BitcoinCmd.
var btcCmdDataName = map[BitcoinCmdData]BitcoinCmdName{
	VersionCmdData: VersionCmd,
}

// btcCmdNameData is a map of BitcoinCmd back to their BitcoinCmdData.
var btcCmdNameData = map[BitcoinCmdName]BitcoinCmdData{
	VersionCmd: VersionCmdData,
}

// BitcoinCmd represents bitcoin command protocol.
type BitcoinCmd struct {
	// HexData represents cmd bytes.
	HexData BitcoinCmdData
	// Name represents cmd name.
	Name BitcoinCmdName
}

// FromHex initialices BitcoinCmd properties from BitcoinCmdData.
func (cmd *BitcoinCmd) FromHex(data []byte) error {
	if len(data) != BitcoinCmdSize {
		return fmt.Errorf("Invalid Bitcoin command size")
	}

	var cmdData [BitcoinCmdSize]byte
	copy(cmdData[:], data)

	name, ok := btcCmdDataName[BitcoinCmdData(cmdData)]
	if !ok {
		return fmt.Errorf("Unknown Bitcoin command")
	}

	cmd.Name = name
	cmd.HexData = BitcoinCmdData(cmdData)
	return nil
}

// FromString initialices BitcoinCmd properties from BitcoinCmdName.
func (cmd *BitcoinCmd) FromString(name string) error {
	data, ok := btcCmdNameData[BitcoinCmdName(name)]
	if !ok {
		return fmt.Errorf("Unknown Bitcoin command")
	}

	cmd.HexData = data
	cmd.Name = BitcoinCmdName(name)
	return nil
}
