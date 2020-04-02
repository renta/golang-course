package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Copy(from string, to string, limit int, offset int) error {

	buf := make([]byte, limit)

	readFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer func() {
		if err := readFile.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = readFile.ReadAt(buf, int64(offset))
	//fmt.Printf("n = %v err = %v b = %v\n", read, err, buf)
	//fmt.Printf("b[:n] = %q\n", buf[:read])
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	writeFile, err := os.OpenFile(to, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_, err = writeFile.Write(buf)
	if err != nil {
		return err
	}
	err = writeFile.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	for i := 0; i < 3; i++ {
		bytesLimit := 1024
		err := Copy("voyna-i-mir.txt", "result.txt", bytesLimit, bytesLimit*i)
		if err != nil {
			fmt.Println(fmt.Errorf("error while copying files with error %v", err.Error()))
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
