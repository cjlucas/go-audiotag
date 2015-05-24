package id3

import (
	"bytes"
	"fmt"
	"time"
)

// Comply with the Getter interface

func (t *ID3) firstInfoString(id ...byte) string {
	for i := range t.Frames {
		f := &t.Frames[i]
		if bytes.Equal(f.Header.FrameId(), id) {
			if f, ok := f.Frame.(*TextInformationFrame); ok {
				return f.Information[0]
			} else {
				return ""
			}
		}
	}

	return ""
}

func (t *ID3) TrackTitle() string {
	return t.firstInfoString('T', 'I', 'T', '2')
}

func (t *ID3) Album() string {
	return t.firstInfoString('T', 'A', 'L', 'B')
}

func (t *ID3) TrackNumber() string {
	return t.firstInfoString('T', 'P', 'O', 'S')
}

func (t *ID3) Date() time.Time {
	var tstr string
	if t.Header.Version[0] == 4 {
		tstr = t.firstInfoString('T', 'D', 'O', 'R')
	} else {
		year := t.firstInfoString('T', 'O', 'R', 'Y')
		mmdd := t.firstInfoString('T', 'D', 'A', 'T')
		tstr = fmt.Sprintf("%s-%s", year, mmdd)
	}

	ti, _ := parseTime(tstr)
	return ti
}
