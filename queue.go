package main

import (
	"io/fs"
	"path/filepath"
)

/*
LOOK FOR ALL MP3 FILES IN A DIRECTORY AND RETURN THEIR PATHS IN A SLICE
*/
func lookForMp3sInDirectory(dir string) []string {
	var ret = []string{}
	// TODO: Learn WalkDir
	/*
	WalkDir is an optimized way to look through a directory recursively and then call a callback function for each file in that directory.
	*/
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(d.Name()) == ".mp3" {
			ret = append(ret, path)
		}
		return nil
	})
	return ret
}