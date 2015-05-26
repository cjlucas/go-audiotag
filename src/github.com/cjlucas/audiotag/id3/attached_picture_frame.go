package id3

import "errors"

type PictureType int

const (
	PictureTypeOther              PictureType = 0x00
	PictureTypeFileIcon           PictureType = 0x01
	PictureTypeOtherFileIcon      PictureType = 0x02
	PictureTypeCoverFront         PictureType = 0x03
	PictureTypeCoverBack          PictureType = 0x04
	PictureTypeLeafletPage        PictureType = 0x05
	PictureTypeMedia              PictureType = 0x06
	PictureTypeLeadArtist         PictureType = 0x07
	PictureTypeArtist             PictureType = 0x08
	PictureTypeConductor          PictureType = 0x09
	PictureTypeBand               PictureType = 0x0A
	PictureTypeComposer           PictureType = 0x0B
	PictureTypeLyricist           PictureType = 0x0C
	PictureTypeRecordingLocation  PictureType = 0x0D
	PictureTypeDuringRecording    PictureType = 0x0E
	PictureTypeDuringPerformance  PictureType = 0x0F
	PictureTypeMovieScreenCapture PictureType = 0x10
	PictureTypeBrightColouredFish PictureType = 0x11
	PictureTypeIllustration       PictureType = 0x12
	PictureTypeBandLogotype       PictureType = 0x13
	PictureTypePublisherLogotype  PictureType = 0x14
)

const minAttachedPictureFrameSize = 4

type AttachedPictureFrame struct {
	TextEncoding Encoding
	MIMEType     string
	PictureType  PictureType
	Description  string
	ImageData    []byte
}

func readUntilNull(buf []byte) ([]byte, error) {
	for i := 0; i < len(buf); i++ {
		if buf[i] == 0 {
			return buf[:i], nil
		}
	}

	return nil, errors.New("reached end")
}

func (f *AttachedPictureFrame) Parse(buf []byte) error {
	if len(buf) < minAttachedPictureFrameSize {
		return errors.New("buffer to small")
	}

	f.TextEncoding = Encoding(buf[0])

	if mime, err := readUntilNull(buf[1:]); err != nil {
		return err
	} else {
		// Inthe event that the MIME media type name is omitted, "image/" will be implied
		if len(mime) == 0 {
			f.MIMEType = "image/"
		} else {
			f.MIMEType = string(mime)
		}

		buf = buf[len(mime)+2:]
	}

	f.PictureType = PictureType(buf[0])

	if desc, err := readUntilNull(buf[1:]); err != nil {
		return err
	} else if str, err := f.TextEncoding.Decode(desc); err != nil {
		return err
	} else {
		f.Description = str
		buf = buf[len(desc)+2:]
	}

	f.ImageData = buf

	return nil
}

// audiotag.Image interface

func (f *AttachedPictureFrame) MIME() string {
	return f.MIMEType
}

func (f *AttachedPictureFrame) Data() []byte {
	return f.ImageData
}
