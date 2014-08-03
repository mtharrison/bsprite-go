package bsprite

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime"
	"path/filepath"
	"strconv"
)

func (sprite Sprite) Headers() map[string]string {
	headers := make(map[string]string)
	headers["X-Metadata-Offset"] = strconv.Itoa(sprite.MetadataOffset)
	headers["Content-type"] = "text/plain"
	return headers
}

func (sprite Sprite) Body() []byte {
	var sum []byte

	for _, img := range sprite.Images {
		for i := 0; i < len(img.data); i++ {
			sum = append(sum, img.data[i])
		}
	}

	for i := 0; i < len(sprite.Metadata); i++ {
		sum = append(sum, sprite.Metadata[i])
	}

	return sum
}

func Make(globs ...string) (err error, sprite Sprite) {

	bytePointer := 0

	files := getFiles(globs)

	for _, file := range files {
		bytes, err := ioutil.ReadFile(file)
		str64 := base64.StdEncoding.EncodeToString(bytes)

		if err != nil {
			log.Fatal(err)
		}

		spriteImage := SpriteImage{
			data:     str64,
			name:     file,
			length:   len(str64),
			offset:   bytePointer,
			mimeType: mime.TypeByExtension(filepath.Ext(file)),
		}

		bytePointer += len(str64)
		sprite.Images = append(sprite.Images, spriteImage)

	}

	sprite.MetadataOffset = bytePointer
	sprite.Metadata = getMetaDataJSON(sprite)

	return
}

func getMetaDataJSON(sprite Sprite) string {
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

	return string(json)
}
