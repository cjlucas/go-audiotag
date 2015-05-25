package audiotag

import "time"

type Getter interface {
	TrackTitle() string
	TrackSubtitle() string
	TrackPosition() string // Format TP/DP
	DiscPosition() string
	DiscSubtitle() string
	AlbumTitle() string
	TrackArtist() string
	TrackArtistSortOrder() string
	AlbumArtist() string
	AlbumArtistSortOrder() string
	Duration() int
	Genre() string
	ReleaseDate() time.Time
	OriginalReleaseDate() time.Time
	//Images() []image.Image // TODO: don't use image.Image
}
