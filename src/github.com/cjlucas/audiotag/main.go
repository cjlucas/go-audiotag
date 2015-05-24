package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cjlucas/audiotag/id3"
)

func parseFile(path string, info os.FileInfo, err error) error {
	if info.IsDir() || filepath.Ext(path) != ".mp3" {
		return nil
	}

	fmt.Println(path)
	if tag, err := id3.Read(path); err != nil {
		panic(err)
	} else {
		fmt.Println(tag.TrackTitle(), tag.Album(), tag.Date())
	}

	return nil
}

func main() {
	arg := os.Args[len(os.Args)-1]
	fi, err := os.Stat(arg)
	if err != nil {
		panic(err)
	}
	if fi.IsDir() {
		filepath.Walk(os.Args[len(os.Args)-1], parseFile)
	} else {
		parseFile(arg, fi, nil)
	}
}
