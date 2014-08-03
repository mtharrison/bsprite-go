package bsprite

type Sprite struct {
	Images         []SpriteImage
	Metadata       string
	MetadataOffset int
}

type SpriteImage struct {
	data     string
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
