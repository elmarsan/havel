package main

import (
	"testing"
)

func TestBitcoinNet(t *testing.T) {
	t.Run("should recognice bitcoin networks", func(t *testing.T) {
		nets := []uint32{
			uint32(MainNet),
			uint32(TestNet),
			uint32(TestNet3),
			uint32(SimNet),
		}

		for _, net := range nets {
			_, err := NewBitcoinNet(net)
			if err != nil {
				t.Errorf("Could not recognice network (0x%x)", net)
			}
		}
	})

	t.Run("should NOT recognice unexisting bitcoin network", func(t *testing.T) {
		var net uint32 = 0xbbb5bff1

		_, err := NewBitcoinNet(net)
		if err == nil {
			t.Errorf("Net (0x%x) should NOT be valid bitcoin network", net)
		}
	})
}

func TestBitcoinCmdFromHex(t *testing.T) {
	t.Run("should return error when unknown command", func(t *testing.T) {
		cmd := &BitcoinCmd{}

		data := BitcoinCmdData{0xaa, 0xbb, 0x72, 0x71, 0x11, 0xf1, 0x65, 0x00, 0x00, 0x00, 0x00, 0x12}
		err := cmd.FromHex(data)
		if err == nil {
			t.Errorf("Command should not be recognized")
		}
	})

	t.Run("should NOT return error", func(t *testing.T) {
		cmd := &BitcoinCmd{}

		err := cmd.FromHex(VersionCmdData)
		if err != nil {
			t.Errorf("Command not recognized")
		}
	})
}

func TestBitcoinCmdFromString(t *testing.T) {
	t.Run("should return error when unknown command", func(t *testing.T) {
		cmd := &BitcoinCmd{}

		err := cmd.FromString("Help")
		if err == nil {
			t.Errorf("Command should not be recognized")
		}
	})

	t.Run("should NOT return error", func(t *testing.T) {
		cmd := &BitcoinCmd{}

		err := cmd.FromString("Version")
		if err != nil {
			t.Errorf("Command not recognized")
		}
	})
}
