package id3

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/cjlucas/audiotag/audiotag"
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

func (t *ID3) TrackSubtitle() string {
	return t.firstInfoString('T', 'I', 'T', '3')
}

func (t *ID3) TrackPosition() string {
	return t.firstInfoString('T', 'R', 'C', 'K')
}

func (t *ID3) DiscPosition() string {
	return t.firstInfoString('T', 'P', 'O', 'S')
}

func (t *ID3) DiscSubtitle() string {
	return t.firstInfoString('T', 'S', 'S', 'T')
}

func (t *ID3) AlbumTitle() string {
	return t.firstInfoString('T', 'A', 'L', 'B')
}

func (t *ID3) TrackArtist() string {
	return t.firstInfoString('T', 'P', 'E', '1')
}

func (t *ID3) TrackArtistSortOrder() string {
	return t.firstInfoString('T', 'S', 'O', 'P')
}

func (t *ID3) AlbumArtist() string {
	return t.firstInfoString('T', 'P', 'E', '2')
}

func (t *ID3) AlbumArtistSortOrder() string {
	return t.firstInfoString('T', 'S', 'O', '2')
}

func (t *ID3) Duration() int {
	i, _ := strconv.Atoi(t.firstInfoString('T', 'L', 'E', 'N'))
	return i
}

func (t *ID3) Genre() string {
	return t.firstInfoString('T', 'C', 'O', 'N')
}

func (t *ID3) ReleaseDate() time.Time {
	var tstr string
	if t.Header.Version[0] == 4 {
		tstr = t.firstInfoString('T', 'D', 'R', 'C')
	} else {
		year := t.firstInfoString('T', 'Y', 'E', 'R')
		mmdd := t.firstInfoString('T', 'D', 'A', 'T')
		tstr = fmt.Sprintf("%s-%s", year, mmdd)
	}

	ti, _ := parseTime(tstr)
	return ti
}

func (t *ID3) OriginalReleaseDate() time.Time {
	var tstr string
	if t.Header.Version[0] == 4 {
		tstr = t.firstInfoString('T', 'D', 'O', 'R')
	} else {
		tstr = t.firstInfoString('T', 'O', 'R', 'Y')
	}

	ti, _ := parseTime(tstr)
	return ti
}

func (t *ID3) Images() []audiotag.Image {
	images := make([]audiotag.Image, 0)
	for _, f := range t.Frames {
		if v, ok := f.Frame.(audiotag.Image); ok {
			images = append(images, v)
		}
	}

	return images
}
