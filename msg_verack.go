package main

import (
	"fmt"
	"io"
)

// https://en.bitcoin.it/wiki/Protocol_documentation#verack
type MsgVerack struct {
	// Header represents msg header.
	Header *MsgHeader
}

// Encode encodes MsgVerack into w.
func (verack *MsgVerack) Encode(w io.Writer) error {
	err := verack.Header.Encode(w)
	if err != nil {
		return fmt.Errorf("Could not encode Headers, cause %s", err.Error())
	}

	return nil
}

// Decode decodes MsgVerack from r.
func (verack *MsgVerack) Decode(r io.Reader) error {
	// Decode headers
	verack.Header = &MsgHeader{}
	err := verack.Header.Decode(r)
	if err != nil {
		return fmt.Errorf("Could not decode Headers, cause %s", err.Error())
	}

	return nil
}
