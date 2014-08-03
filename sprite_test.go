package bsprite

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

const NUM_FILES = 4

type FilesArray [NUM_FILES]string

func TestMake(t *testing.T) {
	_, dirname := makeTempfiles()
	defer destroyTempfiles(dirname)

	// Do test logic

}

func makeTempfiles() (files FilesArray, dirname string) {
	dirname, err := ioutil.TempDir(".", "test")

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < NUM_FILES; i++ {
		file, err := ioutil.TempFile(dirname, "test")
		if err != nil {
			log.Fatal(err)
		}
		files[i] = file.Name()
		file.Write([]byte("Hello, World!"))
	}

	return
}

func destroyTempfiles(dirname string) {
	os.RemoveAll(dirname)
}
