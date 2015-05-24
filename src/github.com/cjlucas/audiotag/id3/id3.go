package id3

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Frame interface {
	Parse([]byte) error
}

type FrameWithHeader struct {
	Header FrameHeader
	Frame  Frame
}

type ID3 struct {
	Header Header
	Frames []FrameWithHeader
}

func frameForId(id []byte) Frame {
	if bytes.Equal(id, []byte{'A', 'P', 'I', 'C'}) {
		return &AttachedPictureFrame{}
	}

	if bytes.Equal(id, []byte{'T', 'X', 'X', 'X'}) {
		return &UserTextInformationFrame{}
	}

	if id[0] == 'T' {
		return &TextInformationFrame{}
	}

	return nil
}

func readUntilId3Identifier(r *bufio.Reader) error {
	var seekbuf [1]byte
	for {
		buf, err := r.Peek(3)
		if err != nil {
			return err
		}

		if bytes.Equal(buf, id3HeaderFileIdentifier[:]) {
			break
		}

		r.Read(seekbuf[:])
	}

	return nil
}

func Read(fpath string) (*ID3, error) {
	id3 := &ID3{}

	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	if err := readUntilId3Identifier(r); err != nil {
		return nil, err
	}

	var id3HeaderBuf [id3HeaderSize]byte
	if _, err := r.Read(id3HeaderBuf[:]); err != nil {
		return nil, err
	} else if err := id3.Header.Parse(id3HeaderBuf[:]); err != nil {
		return nil, err
	}

	buf := make([]byte, id3.Header.Size)
	bytesRead := 0
	for bytesRead < id3.Header.Size {
		if n, err := r.Read(buf[bytesRead:]); err != nil {
			fmt.Println("bytes read:", n)
			return nil, fmt.Errorf("error reading id3: %s\n", err)
		} else {
			bytesRead += n
		}

	}

	bytesRead = 0
	for id3.Header.Size-bytesRead >= frameHeaderSize {
		b := buf[bytesRead:]
		fh := id3v2FrameHeader{UseSynchsafe: id3.Header.Version[0] == 4}
		if err := fh.Parse(b[:frameHeaderSize]); err != nil {
			return nil, err
		}
		bytesRead += frameHeaderSize + fh.Size()

		f := frameForId(fh.FrameId())

		if f == nil {
			continue
		} else if err := f.Parse(b[frameHeaderSize : frameHeaderSize+fh.Size()]); err != nil {
		} else {
			id3.Frames = append(id3.Frames, FrameWithHeader{&fh, f})
		}

	}

	return id3, nil
}
