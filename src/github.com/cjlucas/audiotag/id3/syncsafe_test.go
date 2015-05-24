package id3

import "testing"

func TestConvSyncsafeInt(t *testing.T) {
	cases := []struct {
		input, expected int
	}{
		{0, 0},
		{383, 255},
		{2139062143, 268435455},
	}

	for _, c := range cases {
		actual := ConvSynchsafeInt(c.input)
		if actual != c.expected {
			t.Fatalf("Failed to convert synchsafe int (%d != %d)", actual, c.expected)
		}
	}
}

func BenchmarkConvSynchsafeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvSynchsafeInt(2139062143)
	}
}
