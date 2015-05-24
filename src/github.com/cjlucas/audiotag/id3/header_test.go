package id3

import "testing"

func TestHeaderParse(t *testing.T) {
	cases := []struct {
		in          []byte
		shouldError bool
		expected    Header
	}{
		{
			[]byte{73, 68, 51, 4, 0, 0, 0, 0, 1, 127},
			false,
			Header{
				[3]byte{'I', 'D', '3'},
				[2]byte{4, 0},
				0,
				255,
			},
		},
		{
			[]byte{73, 68, 51, 4, 0, 0, 0, 0, 1},
			true,
			Header{},
		},
		{
			[]byte{74, 68, 51, 4, 0, 0, 0, 0, 1},
			true,
			Header{},
		},
		{
			[]byte{74, 68, 51, 5, 0, 0, 0, 0, 1},
			true,
			Header{},
		},
	}

	h := Header{}
	for _, c := range cases {
		err := h.Parse(c.in)
		if c.shouldError && err == nil {
			t.Fatalf("Parse should have failed, but didnt: %s", c.in)
		}
		if !c.shouldError && err != nil {
			t.Fatalf("Parse should have succeeded, but didnt: %s", c.in)
		}

		if !c.shouldError && h != c.expected {
			t.Fatalf("Header mismatch: %#v != %#v", h, c.expected)
		}
	}
}
