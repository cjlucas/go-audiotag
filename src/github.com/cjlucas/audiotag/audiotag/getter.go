package audiotag

import (
	"image"
	"time"
)

type Getter interface {
	Title() string
	Subtitle() string
	TrackNumber() string // Format TP/DP
	Album() string
	Artist() string
	AlbumArtist() string
	Date() time.Time
	Images() []image.Image
}
