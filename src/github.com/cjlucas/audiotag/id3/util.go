package id3

import (
	"bytes"
	"errors"
	"io"
	"time"
)

var timestampFormats = []string{
	"2006-01-02T15:04:05",
	"2006-01-02T15:04",
	"2006-01-02T15",
	"2006-01-02",
	"2006-01",
	"2006",
}

func parseTime(timeStr string) (time.Time, error) {
	for i := range timestampFormats {
		t, err := time.Parse(timestampFormats[i], timeStr)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, errors.New("invalid time")
}

func readUntilTerminator(term []byte, buf []byte) ([]byte, error) {
	for i := 0; i+len(term)-1 < len(buf); i += len(term) {
		if bytes.Equal(term, buf[i:i+len(term)]) {
			return buf[:i], nil
		}
	}

	return nil, io.EOF
}
