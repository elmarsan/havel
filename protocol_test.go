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

func TestBitcoinCmd(t *testing.T) {
	t.Run("should recognice bitcoin command", func(t *testing.T) {
		cmds := []string{VersionCmdHex}

		for _, cmd := range cmds {
			_, err := NewBitcoinCmd(cmd)
			if err != nil {
				t.Errorf("Could not recognice command (%s)", cmd)
			}
		}
	})

	t.Run("should NOT recognice unexisting commands", func(t *testing.T) {
		cmd := "0xf4f5f6abae1298"

		_, err := NewBitcoinCmd(cmd)
		if err == nil {
			t.Errorf("Cmd (%s) should NOT be valid bitcoin command", cmd)
		}
	})
}
