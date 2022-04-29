package main

// 在文件中查找位置
import (
	"errors"
	"fmt"
	"os"
)

const lineLength = 25

func main() {

	f, e := os.OpenFile("flatfile.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	fmt.Println(readRecords(2, "last", f))
	if err := writeRecord(2, "first", "Radomir", f); err != nil {
		panic(err)
	}
	fmt.Println(readRecords(2, "first", f))
	if err := writeRecord(10, "first", "Andrew", f); err != nil {
		panic(err)
	}
	fmt.Println(readRecords(10, "first", f))
	fmt.Println(readLine(2, f))
}

func readLine(line int, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.Seek(int64(line*lineLength), 0)
	_, err := f.Read(lineBuffer)
	return string(lineBuffer), err
}

func writeRecord(line int, column, dataStr string, f *os.File) error {
	defineLen := 10
	position := int64(line * lineLength)
	switch column {
	case "id":
		defineLen = 4
	case "first":
		position += 4
	case "last":
		position += 14
	default:
		return errors.New("Column not defined")
	}

	if len([]byte(dataStr)) > defineLen {
		return fmt.Errorf("Maximum length for %s is %d", column, defineLen)
	}

	data := make([]byte, defineLen)
	for i := range data {
		data[i] = '.'
	}
	copy(data, []byte(dataStr))
	_, err := f.WriteAt(data, position)
	return err
}

func readRecords(line int, column string, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.ReadAt(lineBuffer, int64(line*lineLength))
	var retVal string
	switch column {
	case "id":
		return string(lineBuffer[:3]), nil
	case "first":
		return string(lineBuffer[4:13]), nil
	case "last":
		return string(lineBuffer[14:23]), nil
	}

	return retVal, errors.New("Column not defined")
}
