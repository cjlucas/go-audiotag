package id3

import (
	"encoding/binary"
	"errors"
)

const id3HeaderSize = 10

var id3HeaderFileIdentifier = [3]byte{'I', 'D', '3'}

type Header struct {
	FileIdentifier [3]byte
	Version        [2]byte
	Flags          byte
	Size           int
}

func (h *Header) Parse(buf []byte) error {
	if len(buf) < id3HeaderSize {
		return errors.New("received buffer of unexpected size")
	}

	for i := 0; i < len(h.FileIdentifier); i++ {
		h.FileIdentifier[i] = buf[i]
	}

	copy(h.FileIdentifier[:], buf[0:3])
	copy(h.Version[:], buf[3:5])

	if h.FileIdentifier != id3HeaderFileIdentifier {
		return errors.New("unexpected file identifier")
	}

	if h.Version[0] != 4 && h.Version[0] != 3 {
		return errors.New("unexpected version")
	}

	h.Flags = buf[6]
	h.Size = int(binary.BigEndian.Uint32(buf[6:10]))
	h.Size = ConvSynchsafeInt(h.Size)

	return nil
}
