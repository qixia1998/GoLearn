package main

// 有效解析大型 XML 文件
import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
}

func main() {

	f, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)

	// Read the book one by one
	books := make([]Book, 0)
	for {
		tok, _ := decoder.Token()
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "book" {
				// Decode the element to struct
				var b Book
				decoder.DecodeElement(&b, &tp)
				books = append(books, b)
			}
		}
	}
	fmt.Println(books)
}
