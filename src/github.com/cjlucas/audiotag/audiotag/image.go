package audiotag

type Image interface {
	MIME() string
	Data() []byte
}
