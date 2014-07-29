package bsprite

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime"
	"path/filepath"
	"strconv"
)

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
	var metadata SpriteMetadata

	for _, i := range sprite.Images {
		m := SpriteImageMetaData{
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