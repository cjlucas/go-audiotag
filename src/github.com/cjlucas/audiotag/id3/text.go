package id3

import (
	"bytes"
	"errors"
	"unicode/utf16"
)

type Encoding int

const (
	ISO88591 Encoding = 0x00
	UTF16    Encoding = 0x01
	UTF16BE  Encoding = 0x02
	UTF8     Encoding = 0x03
)

type decoder func([]byte) (string, error)

var utf16BOM = []byte{0xFF, 0xFE}

var encodingTerminatorMap = map[Encoding][]byte{
	ISO88591: []byte{0x00},
	UTF16:    []byte{0x00, 0x00},
	UTF16BE:  []byte{0x00, 0x00},
	UTF8:     []byte{0x00},
}

var encodingDecoderMap = map[Encoding]decoder{
	ISO88591: decodeISO88591,
	UTF16:    decodeUTF16,
	UTF16BE:  decodeUTF16BE,
	UTF8:     decodeUTF8,
}

func reachedEnd(term []byte, data []byte) bool {
	if len(data) < len(term) {
		return true
	}

	for i := 0; i < len(term); i++ {
		if term[i] != data[i] {
			return false
		}
	}

	return true
}

func decodeISO88591(data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}

	term := encodingTerminatorMap[ISO88591]

	end := len(data)
	for i := 0; i < len(data); i++ {
		if reachedEnd(term, data[i:]) {
			end = i
			break
		}
	}

	return string(data[:end]), nil
}

func decodeUTF16(data []byte) (string, error) {
	if len(data) < 2 {
		return "", nil
	}

	// Skip BOM if present
	if bytes.Equal(data[:2], utf16BOM) {
		data = data[2:]
	}

	if len(data)%2 != 0 {
		return "", errors.New("received non-even number of bytes")
	}

	points := make([]uint16, 0, len(data)/2)
	term := encodingTerminatorMap[UTF16]
	for i := 0; i+1 < len(data) && !reachedEnd(term, data[i:]); i += 2 {
		msb := uint16(data[i])
		lsb := uint16(data[i+1])
		points = append(points, (lsb<<8)|msb)
	}

	return string(utf16.Decode(points)), nil
}
func decodeUTF16BE(data []byte) (string, error) {
	// swap code points to LE
	var temp byte
	for i := 0; i+1 < len(data); i += 2 {
		temp = data[i]
		data[i] = data[i+1]
		data[i+1] = temp
	}

	return decodeUTF16(data)
}
func decodeUTF8(data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}

	term := encodingTerminatorMap[UTF8]
	end := len(data)
	for i := 0; i < len(data); i++ {
		if reachedEnd(term, data[i:]) {
			end = i
			break
		}
	}

	return string(data[:end]), nil
}

func (enc Encoding) Decode(data []byte) (string, error) {
	return encodingDecoderMap[enc](data)
}

func (enc Encoding) Terminator() []byte {
	return encodingTerminatorMap[enc]
}
