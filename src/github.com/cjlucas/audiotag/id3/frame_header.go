package id3

import (
	"encoding/binary"
	"errors"
)

const frameHeaderSize = 10

type FrameHeaderFlags [2]byte

type FrameHeader interface {
	FrameId() []byte
	Size() int
	Flags() FrameHeaderFlags
}

type id3v2FrameHeader struct {
	UseSynchsafe bool
	frameId      [4]byte
	size         int // raw
	flags        FrameHeaderFlags
}

func (h *id3v2FrameHeader) FrameId() []byte {
	return h.frameId[:]
}

func (h *id3v2FrameHeader) Size() int {
	if h.UseSynchsafe {
		return ConvSynchsafeInt(h.size)
	} else {
		return h.size
	}
}

func (h *id3v2FrameHeader) Flags() FrameHeaderFlags {
	return h.flags
}

func (h *id3v2FrameHeader) Parse(buf []byte) error {
	if len(buf) < 10 {
		return errors.New("buffer too small")
	}

	copy(h.frameId[:], buf[0:4])
	h.size = int(binary.BigEndian.Uint32(buf[4:8]))
	copy(h.flags[:], buf[8:10])

	return nil
}
