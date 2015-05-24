package id3

import "testing"

func TestTextInformationFrameParse(t *testing.T) {
	cases := []struct {
		in          []byte
		expected    TextInformationFrame
		shouldError bool
	}{
		{
			// single information string
			[]byte{byte(ISO88591), 'h', 'i', 't', 'h', 'e', 'r', 'e'},
			TextInformationFrame{
				TextEncoding: ISO88591,
				Information:  []string{"hithere"},
			},
			false,
		},
		{
			// multiple information strings
			[]byte{byte(ISO88591), 'h', 'i', 't', 'h', 'e', 'r', 'e', 0, 'w', 'h', 'a', 't'},
			TextInformationFrame{
				TextEncoding: ISO88591,
				Information:  []string{"hithere", "what"},
			},
			false,
		},
		{
			// UTF16: multiple information strings
			[]byte{byte(UTF16), 'h', 0, 'i', 0, 0, 0, 'y', 0, 'o', 0, 'u', 0},
			TextInformationFrame{
				TextEncoding: UTF16,
				Information:  []string{"hi", "you"},
			},
			false,
		},
		{
			[]byte{},
			TextInformationFrame{},
			true,
		},
	}

	for i, c := range cases {
		f := TextInformationFrame{}
		err := f.Parse(c.in)

		if c.shouldError && err != nil {
			continue
		}

		if c.shouldError && err == nil {
			t.Fatal("Parsing unexpectedly succeeeded for case:", i)
		} else if !c.shouldError && err != nil {
			t.Fatalf("Parsing unexpectedly failed for case %d (err: %s)", i, err)
		}

		if f.TextEncoding != c.expected.TextEncoding {
			t.Fatalf("Text encoding mismatch for case %d: (%d != %d)",
				i,
				f.TextEncoding,
				c.expected.TextEncoding,
			)
		}

		if len(f.Information) != len(c.expected.Information) {
			t.Fatalf("Information count mismatch for case %d: (%d != %d)",
				i,
				len(f.Information),
				len(c.expected.Information),
			)
		}

		for i := range f.Information {
			if f.Information[i] != c.expected.Information[i] {
				t.Fatalf("Information string mismatch for case %d: (%s != %s)",
					i,
					f.Information[i],
					c.expected.Information[i],
				)
			}
		}
	}
}
