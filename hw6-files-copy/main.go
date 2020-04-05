package main

import (
	"bytes"
	"fmt"
	"github.com/renta/golang-course/hw6-files-copy/internal/file"
	"io/ioutil"
)

func main() {
	for i := 0; i < 3; i++ {
		bytesLimit := 1024
		writtenBytes, err := file.Copy("voyna-i-mir.txt", "result.txt", bytesLimit*i, bytesLimit)
		if err != nil {
			fmt.Println(fmt.Errorf("error while copying files with error %v", err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("bytes were written %d", writtenBytes))
		}
	}

	resultFile, err := ioutil.ReadFile("result.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can not open result.txt"))
	}

	referenceFile, err := ioutil.ReadFile("result_reference.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("can not open result.txt"))
	}

	result := bytes.Equal(resultFile, referenceFile)
	if result {
		fmt.Println("Copy function works as expected")
	} else {
		fmt.Println(fmt.Errorf("result.txt and result_reference.txt have not equal content, Copy function failed"))
	}
}
