package id3

import "testing"

func TestUserTextInformationFrameParse(t *testing.T) {
	cases := []struct {
		in          []byte
		expected    UserTextInformationFrame
		shouldError bool
	}{
		{
			[]byte{0, 'h', 'i', 't', 'h', 'e', 'r', 'e', 0, 's', 'u', 'p'},
			UserTextInformationFrame{
				TextEncoding: ISO88591,
				Description:  "hithere",
				Value:        "sup",
			},
			false,
		},
	}

	for _, c := range cases {
		f := UserTextInformationFrame{}
		err := f.Parse(c.in)
		if c.shouldError && err != nil {
			continue
		}

		if c.shouldError && err == nil {
			t.Fatal("expected error")
		} else if !c.shouldError && err != nil {
			t.Fatal("unexpected error")
		}

		if f.Description != c.expected.Description {
			t.Fatalf("Description mismatch (%s != %s)",
				f.Description,
				c.expected.Description,
			)
		}

		if f.Value != c.expected.Value {
			t.Fatalf("Value mismatch (%s != %s)",
				f.Value,
				c.expected.Value,
			)
		}
	}
}
