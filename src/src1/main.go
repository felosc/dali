package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {

	type walkFunc func(path string, info fs.FileInfo, err error) error
	count := 0
	countDir := 0
	var pathEmailToWalk = "/home/vboxuser/Downloads/enron_mail_20110402/maildir"
	dirFiles := os.DirFS(pathEmailToWalk)
	fs.WalkDir(dirFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		if d.Type().IsDir() {
			countDir++
			println(path, " es carpeta N°: ", countDir, "\n")
		} else {
			fmt.Println("Archivo: ", path)
			count++
		}

		return nil
	})
}
