package file

import (
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	FileCopyFrom      string = "test.txt"
	FileCopyToPattern string = "test_%s_result.txt"
)

func TestMain(m *testing.M) {
	err := createFileCopyFrom()
	if err != nil {
		log.Fatal(err)
	}
	code := m.Run()
	deleteFile(FileCopyFrom)
	os.Exit(code)
}

func TestCopyWhole(t *testing.T) {
	fileSize, err := GetFileSize(FileCopyFrom)
	if err != nil {
		t.Errorf(err.Error())
	}

	fileNameForCopied := copiedFileNameFromPattern("copy_whole")
	writtenBytes, err := Copy(FileCopyFrom, fileNameForCopied, 0, fileSize)
	if err != nil {
		t.Errorf(err.Error())
	}

	if writtenBytes != fileSize || !isFilesAreEqual(t, FileCopyFrom, fileNameForCopied) {
		t.Errorf("result file is not equal to givven one")
	} else {
		deleteFile(fileNameForCopied)
	}
}

func TestCopyPartFromBeginning(t *testing.T) {
	bytesToCopy := 10
	fileNameForCopied := copiedFileNameFromPattern("copy_part_from_beginning")
	writtenBytes, err := Copy(FileCopyFrom, fileNameForCopied, 0, bytesToCopy)
	if err != nil {
		t.Errorf(err.Error())
	}

	if writtenBytes != bytesToCopy || !fileContentEqualToString(t, fileNameForCopied, "1_2_3_4_5_") {
		t.Errorf("result file is not equal to givven one")
	} else {
		deleteFile(fileNameForCopied)
	}
}

func TestCopyPartWithOffset(t *testing.T) {
	bytesToCopy := 10
	fileNameForCopied := copiedFileNameFromPattern("copy_part_with_offset")
	writtenBytes, err := Copy(FileCopyFrom, fileNameForCopied, 10, bytesToCopy)
	if err != nil {
		t.Errorf(err.Error())
	}

	if writtenBytes != bytesToCopy || !fileContentEqualToString(t, fileNameForCopied, "6_7_8_9_10") {
		t.Errorf("result file is not equal to givven one")
	} else {
		deleteFile(fileNameForCopied)
	}
}

func TestCopyWithOffsetToEnd(t *testing.T) {
	fileNameForCopied := copiedFileNameFromPattern("copy_part_with_offset")
	bytesOffset := 260
	writtenBytes, err := Copy(FileCopyFrom, fileNameForCopied, bytesOffset, 0)
	if err != nil {
		t.Errorf(err.Error())
	}

	copiedFileSize, err := GetFileSize(FileCopyFrom)
	if err != nil {
		t.Errorf(err.Error())
	}
	bytesToCopy := copiedFileSize - bytesOffset

	if writtenBytes != bytesToCopy || !fileContentEqualToString(t, fileNameForCopied, "_91_92_93_94_95_96_97_98_99_100") {
		t.Errorf("result file is not equal to givven one")
	} else {
		deleteFile(fileNameForCopied)
	}
}

func TestCopyNoOffsetNoLimit(t *testing.T) {
	fileSize, err := GetFileSize(FileCopyFrom)
	if err != nil {
		t.Errorf(err.Error())
	}

	fileNameForCopied := copiedFileNameFromPattern("copy_no_offset_no_limit")
	writtenBytes, err := Copy(FileCopyFrom, fileNameForCopied, 0, 0)
	if err != nil {
		t.Errorf(err.Error())
	}

	if writtenBytes != fileSize || !isFilesAreEqual(t, FileCopyFrom, fileNameForCopied) {
		t.Errorf("result file is not equal to givven one")
	} else {
		deleteFile(fileNameForCopied)
	}
}

func TestCopyWithOffsetAndLimitHigherThanEOF(t *testing.T) {
	fileNameForCopied := copiedFileNameFromPattern("copy_part_with_offset_and_lim_eof")
	bytesOffset := 285
	writtenBytes, err := Copy(FileCopyFrom, fileNameForCopied, bytesOffset, 50)
	if err != nil {
		t.Errorf(err.Error())
	}

	copiedFileSize, err := GetFileSize(FileCopyFrom)
	if err != nil {
		t.Errorf(err.Error())
	}
	bytesToCopy := copiedFileSize - bytesOffset

	if writtenBytes != bytesToCopy || !fileContentEqualToString(t, fileNameForCopied, "99_100") {
		t.Errorf("result file is not equal to givven one")
	} else {
		deleteFile(fileNameForCopied)
	}
}

func TestWrongFuncParams(t *testing.T) {
	tests := []struct {
		name          string
		from          string
		to            string
		offset        int
		limit         int
		copiedBytes   int
		resultMessage error
	}{
		{
			"empty from", "", "somefile.txt", 0, 0, 0, CopyMisconfiguration,
		},
		{
			"empty to", "somefile.txt", "", 0, 0, 0, CopyMisconfiguration,
		},
		{
			"negative offset", "somefile.txt", "", -1, 0, 0, CopyMisconfiguration,
		},
		{
			"negative limit", "somefile.txt", "", 0, 1, 0, CopyMisconfiguration,
		},
		{
			"empty and negative all", "", "", -1, -1, 0, CopyMisconfiguration,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copiedBytes, err := Copy(tt.from, tt.to, tt.offset, tt.limit)
			if copiedBytes > 0 || err != CopyMisconfiguration {
				t.Errorf("[Copy()] failed test %s", tt.name)
			}
		})
	}
}

func createFileCopyFrom() error {
	var fileContent string
	for i := 1; i <= 100; i++ {
		delimiter := "_"
		if i == 100 {
			delimiter = ""
		}
		fileContent += strconv.Itoa(i) + delimiter
	}

	writeFile, err := os.OpenFile(FileCopyFrom, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	_, err = writeFile.Write([]byte(fileContent))
	if err != nil {
		return err
	}

	err = writeFile.Close()
	if err != nil {
		return err
	}

	return nil
}

func copiedFileNameFromPattern(replTo string) string {
	return strings.Replace(FileCopyToPattern, "%s", replTo, 1)
}
