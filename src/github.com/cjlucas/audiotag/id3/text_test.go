package id3

import "testing"

func TestDecode(t *testing.T) {
	cases := []struct {
		enc         Encoding
		in          []byte
		out         string
		shouldError bool
	}{
		{
			ISO88591,
			[]byte{103, 111, 112, 104, 101, 114, 115},
			"gophers",
			false,
		},
		{
			ISO88591,
			[]byte{103, 111, 112, 104, 101, 114, 115, 0},
			"gophers",
			false,
		},
		{
			UTF8,
			[]byte{72, 101, 108, 108, 111, 44, 32, 228, 184, 150, 231, 149, 140},
			"Hello, 世界",
			false,
		},
		{
			UTF8,
			[]byte{72, 101, 108, 108, 111, 44, 32, 228, 184, 150, 231, 149, 140, 0},
			"Hello, 世界",
			false,
		},
		{
			UTF16,
			[]byte{72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 22, 78, 76, 117},
			"Hello, 世界",
			false,
		},
		{
			UTF16,
			[]byte{72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 22, 78, 76, 117, 0, 0},
			"Hello, 世界",
			false,
		},
		{
			UTF16,
			[]byte{255, 254, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 0, 22, 78, 76, 117, 0, 0},
			"Hello, 世界",
			false,
		},
		{
			UTF16BE,
			[]byte{0, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 78, 22, 117, 76},
			"Hello, 世界",
			false,
		},
		{
			UTF16BE,
			[]byte{0, 72, 0, 101, 0, 108, 0, 108, 0, 111, 0, 44, 0, 32, 78, 22, 117, 76, 0, 0},
			"Hello, 世界",
			false,
		},
		{ISO88591, []byte{}, "", false},
		{UTF16, []byte{}, "", false},
		{UTF16BE, []byte{}, "", false},
		{UTF8, []byte{}, "", false},
		{ISO88591, []byte{0, 40}, "", false},
		{UTF16, []byte{0, 0, 40, 41}, "", false},
		{UTF16BE, []byte{0, 0, 40, 41}, "", false},
		{UTF8, []byte{0, 40}, "", false},
	}

	for _, c := range cases {
		actual, err := c.enc.Decode(c.in)
		if c.shouldError && err == nil {
			t.Fatalf("Expected error, but decoding succeeded: %s", c.in)
		}

		if !c.shouldError && err != nil {
			t.Fatalf("Expected no error, but decoding failed: %s", c.in)
		}

		if !c.shouldError && c.out != actual {
			t.Fatalf("Decoding mismatch: %s != %s (in: %v)", actual, c.out, c.in)
		}
	}
}
