package msg

import (
	"fmt"
	"io"
)

// https://en.bitcoin.it/wiki/Protocol_documentation#verack
type Verack struct {
	// Header represents msg header.
	Header *Header
}

// Encode encodes Verack into w.
func (verack *Verack) Encode(w io.Writer) error {
	err := verack.Header.Encode(w)
	if err != nil {
		return fmt.Errorf("Could not encode Headers, cause %s", err.Error())
	}

	return nil
}

// Decode decodes Verack from r.
func (verack *Verack) Decode(r io.Reader) error {
	// Decode headers
	verack.Header = &Header{}
	err := verack.Header.Decode(r)
	if err != nil {
		return fmt.Errorf("Could not decode Headers, cause %s", err.Error())
	}

	return nil
}
