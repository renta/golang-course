package main

import (
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/renta/golang-course/hw6-files-copy/internal/file"
	"io"
	"os"
)

func main() {
	var from, to string
	var offset, limit int

	flag.StringVar(&from, "from", "", "which file to copy")
	flag.StringVar(&to, "to", "", "where file to copy")
	flag.IntVar(&offset, "offset", 0, "offset in file to copy from byte")
	flag.IntVar(&limit, "limit", 0, "limit of bytes to copy from file")

	flag.Parse()

	if from == "" {
		panic("filename to copy should be defined")
	}
	if to == "" {
		panic("filename where to copy should be defined")
	}

	const FileChunksNumber = 10

	fileSize, err := file.GetFileSize(from)
	if err != nil {
		fmt.Println(fmt.Errorf("error while copying files with error %v", err.Error()))
	}

	if offset > fileSize {
		fmt.Println(fmt.Errorf("offset is not in the filesize"))
	}

	var bytesCopyIter, barSize int

	if limit == 0 {
		bytesCopyIter = (fileSize - offset) / FileChunksNumber
		barSize = fileSize - offset
	} else {
		bytesCopyIter = limit / FileChunksNumber
		barSize = limit
	}

	bar := pb.StartNew(barSize)
	bar.SetWriter(os.Stdout)

	for {
		copiedBytes, err := file.Copy(from, to, offset, bytesCopyIter)
		if err != nil && err != io.EOF {
			fmt.Println(fmt.Errorf("error while copying files with error %v", err.Error()))
		}

		offset += copiedBytes
		bar.Add(copiedBytes)

		if offset == fileSize {
			break
		}
	}

	bar.Finish()

	fmt.Println("")
}
