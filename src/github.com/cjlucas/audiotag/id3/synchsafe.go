package id3

func ConvSynchsafeInt(i int) int {
	return ((i & 0x7F000000) >> 3) | ((i & 0x7F0000) >> 2) | ((i & 0x7F00) >> 1) | (i & 0x7F)
}
