package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name 	string
	Surname string
	Number 	string
	LastAccess string
}

var myData = []Record{}

func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 一次性读取所有 CSV 文件
	// 行数据类型是 [][]string
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	// 更改默认字段分隔符为tab
	csvwriter.Comma = '\t'
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("csvData input output!")
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	lines, err := readCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// CSV 数据按列读取， 每行是一个切片
	for _, line := range lines {
		temp := Record{
			Name: line[0],
			Surname: line[1],
			Number: line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		fmt.Println(temp)
	}

	err = saveCSVFile(output)
	if err != nil {
		fmt.Println(err)
		return
	}
}
