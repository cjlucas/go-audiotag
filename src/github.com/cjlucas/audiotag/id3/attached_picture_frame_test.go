package id3

import (
	"bytes"
	"testing"
)

func TestAttachedPictureFrameParse(t *testing.T) {
	cases := []struct {
		description string
		in          []byte
		expected    AttachedPictureFrame
		shouldError bool
	}{
		{
			"valid; exact size buf; mimetype given; desc given",
			[]byte{
				0,
				'i', 'm', 'a', 'g', 'e', '/', 'p', 'n', 'g', 0,
				0x03,
				't', 'h', 'e', ' ', 'c', 'o', 'v', 'e', 'r', 0,
				0, 1, 2, 3, 4, 5,
			},
			AttachedPictureFrame{
				TextEncoding: ISO88591,
				MIMEType:     "image/png",
				PictureType:  PictureTypeCoverFront,
				Description:  "the cover",
				Data:         []byte{0, 1, 2, 3, 4, 5},
			},
			false,
		},
		{
			"valid; mimetype omitted; description ommited",
			[]byte{
				0,
				0,
				0x03,
				0,
				0, 1, 2, 3, 4, 5,
			},
			AttachedPictureFrame{
				TextEncoding: ISO88591,
				MIMEType:     "image/",
				PictureType:  PictureTypeCoverFront,
				Description:  "",
				Data:         []byte{0, 1, 2, 3, 4, 5},
			},
			false,
		},
		{
			"invalid; empty buf",
			[]byte{},
			AttachedPictureFrame{},
			true,
		},
		{
			"invalid; buf too small",
			[]byte{0, 0},
			AttachedPictureFrame{},
			true,
		},
	}

	for _, c := range cases {
		f := AttachedPictureFrame{}
		err := f.Parse(c.in)

		if c.shouldError && err != nil {
			continue
		}

		if c.shouldError && err == nil {
			t.Fatal("Parsing unexpectedly succeeeded for case:", c.description)
		} else if !c.shouldError && err != nil {
			t.Fatalf("Parsing unexpectedly failed for case \"%s\" (err: %s)", c.description, err)
		}

		if f.TextEncoding != c.expected.TextEncoding {
			t.Fatalf("Text encoding mismatch for (case: %s) (%d != %d)",
				c.description,
				f.TextEncoding,
				c.expected.TextEncoding,
			)
		}

		if f.MIMEType != c.expected.MIMEType {
			t.Fatalf("Mimetype mismatch (case: %s) (%s != %s)",
				c.description,
				f.MIMEType,
				c.expected.MIMEType,
			)
		}

		if f.PictureType != c.expected.PictureType {
			t.Fatalf("Picture type mismatch (case: %s) (%d != %d)",
				c.description,
				f.PictureType,
				c.expected.PictureType,
			)

		}

		if f.Description != c.expected.Description {
			t.Fatalf("Picture type mismatch (case: %s) (%s != %s)",
				c.description,
				f.Description,
				c.expected.Description,
			)
		}

		if !bytes.Equal(f.Data, c.expected.Data) {
			t.Fatalf("Data mismatch (case: %s) (%v != %v)",
				c.description,
				f.Data,
				c.expected.Data,
			)
		}
	}
}
