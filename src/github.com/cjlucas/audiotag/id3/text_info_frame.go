package id3

import "errors"

type TextInformationFrame struct {
	TextEncoding Encoding
	Information  []string
}

const minTextInformationFrameSize = 2

func splitTerminator(term []byte, data []byte) [][]byte {
	out := make([][]byte, 0)

	for {
		if str, err := readUntilTerminator(term, data); err != nil {
			break
		} else {
			out = append(out, str)
			data = data[len(str)+len(term):]
		}
	}

	out = append(out, data)
	return out
}

func (f *TextInformationFrame) Parse(buf []byte) error {
	if len(buf) < minTextInformationFrameSize {
		return errors.New("buffer too small")
	}

	f.TextEncoding = Encoding(buf[0])

	info := splitTerminator(f.TextEncoding.Terminator(), buf[1:])
	f.Information = make([]string, len(info))
	for i := range info {
		str, err := f.TextEncoding.Decode(info[i])
		if err != nil {
			return errors.New("error decoding text")
		}

		f.Information[i] = str
	}

	return nil
}
