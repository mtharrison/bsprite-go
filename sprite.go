package bsprite

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime"
	"path/filepath"
	"strconv"
)

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

type MetadataSpec struct {
	Name     string
	MimeType string
	Offset   int
	Length   int
}

type Metadata []MetadataSpec

func (sprite Sprite) Headers() map[string]string {

	headers := make(map[string]string)

	headers["X-Metadata-Length"] = strconv.Itoa(len(sprite.Metadata))

	return headers
}

func (sprite Sprite) Body() []byte {

	var sum []byte

	for i := 0; i < len(sprite.Metadata); i++ {
		sum = append(sum, sprite.Metadata[i])
	}

	for _, img := range sprite.Images {

		for i := 0; i < len(img.data); i++ {
			sum = append(sum, img.data[i])
		}

	}

	return sum
}

func Make(globs ...string) (err error, sprite Sprite) {

	bytePointer := 0

	files := getFiles(globs)

	for _, file := range files {

		bytes, err := ioutil.ReadFile(file)

		if err != nil {
			log.Fatal(err)
		}

		spriteImage := SpriteImage{
			data:     bytes,
			name:     file,
			length:   len(bytes),
			offset:   bytePointer,
			mimeType: mime.TypeByExtension(filepath.Ext(file)),
		}

		bytePointer += len(bytes) + 1

		sprite.Images = append(sprite.Images, spriteImage)

	}

	sprite.Metadata = getMetaDataJSON(sprite)

	return
}

func getMetaDataJSON(sprite Sprite) []byte {
	var metadata Metadata

	for _, i := range sprite.Images {
		m := MetadataSpec{
			Name:     i.name,
			MimeType: i.mimeType,
			Offset:   i.offset,
			Length:   i.length,
		}

		metadata = append(metadata, m)
	}

	json, _ := json.Marshal(metadata)

	return json
}

// func GetAll(globs ...string) (err error, headers map[string]string, body []byte) {

// }

func getFiles(globs []string) []string {

	var files []string

	for _, glob := range globs {

		matches, _ := filepath.Glob(glob)

		for _, match := range matches {
			files = appendIfMissing(files, match)
		}

	}

	return files
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}
