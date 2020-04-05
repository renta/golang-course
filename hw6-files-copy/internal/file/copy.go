package file

import (
	"errors"
	"io"
	"os"
)

var CopyMisconfiguration = errors.New("wrong parameter for Copy function")

func Copy(from string, to string, offset int, limit int) (int, error) {
	if from == "" || to == "" || offset < 0 || limit < 0 {
		return 0, CopyMisconfiguration
	}

	var written int
	readFile, err := os.Open(from)
	if err != nil {
		return written, err
	}
	defer func() {
		if err := readFile.Close(); err != nil {
			panic(err)
		}
	}()

	if limit == 0 {
		fi, err := readFile.Stat()
		if err != nil {
			return written, err
		}
		limit = int(fi.Size()) - offset
	}

	buf := make([]byte, limit)

	_, err = readFile.ReadAt(buf, int64(offset))
	if err == io.EOF {
		return written, nil
	}
	if err != nil {
		return written, err
	}

	writeFile, err := os.OpenFile(to, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return written, err
	}

	written, err = writeFile.Write(buf)
	if err != nil {
		return written, err
	}

	err = writeFile.Close()
	if err != nil {
		return written, err
	}

	return written, nil
}
