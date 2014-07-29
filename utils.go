package bsprite

import "path/filepath"

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
