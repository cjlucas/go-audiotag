package id3

import "errors"

type UserTextInformationFrame struct {
	TextEncoding Encoding
	Description  string
	Value        string
}

const minUserTextInformationFrameSize = 4

func (f *UserTextInformationFrame) Parse(buf []byte) error {
	if len(buf) < minTextInformationFrameSize {
		return errors.New("buffer too small")
	}

	f.TextEncoding = Encoding(buf[0])

	term := f.TextEncoding.Terminator()
	if desc, err := readUntilTerminator(term, buf[1:]); err != nil {
		return err
	} else if str, err := f.TextEncoding.Decode(desc); err != nil {
		return err
	} else {
		f.Description = str
		buf = buf[1+len(desc)+len(term):]
	}

	if str, err := f.TextEncoding.Decode(buf); err != nil {
		return err
	} else {
		f.Value = str
	}

	return nil
}
