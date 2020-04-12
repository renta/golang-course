package file

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func GetFileSize(fileName string) (int, error) {
	fileInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return 0, CopyFileIsNotExists
	}
	if err != nil {
		return 0, err
	}
	if fileInfo.IsDir() {
		return 0, CopyFileIsDirectory
	}

	return int(fileInfo.Size()), nil
}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func fileContentEqualToString(t *testing.T, file string, content string) bool {
	resultFile, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("can not open %s", file)
	}

	return bytes.Equal(resultFile, []byte(content))
}

func isFilesAreEqual(t *testing.T, first string, second string) bool {
	resultFile, err := ioutil.ReadFile(first)
	if err != nil {
		t.Errorf("can not open %s", first)
	}

	referenceFile, err := ioutil.ReadFile(second)
	if err != nil {
		t.Errorf("can not open %s", second)
	}

	return bytes.Equal(resultFile, referenceFile)
}
