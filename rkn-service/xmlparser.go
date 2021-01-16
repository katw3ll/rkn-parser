package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Поля, которые необходимо парсить
type Record struct {
	INN string `xml:"inn" bson:"inn"`
}

func Parsing() {
	f, err := os.Open("/tmp/data.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)

	for {
		tok, err := decoder.Token()
		if err != nil {
			panic(err)
		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "record" {
				// Декодирование элемента в структуру
				var b Record
				decoder.DecodeElement(&b, &tp)
				fmt.Println(b)
			}
		}
	}
}
