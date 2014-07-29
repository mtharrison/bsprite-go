package bsprite

type Sprite struct {
	Images   []SpriteImage
	Metadata []byte
}

type SpriteImage struct {
	data     []byte
	name     string
	offset   int
	length   int
	mimeType string
}

type SpriteImageMetaData struct {
	Name     string
	MimeType string
	Offset   int
	Length   int
}

type SpriteMetadata []SpriteImageMetaData
