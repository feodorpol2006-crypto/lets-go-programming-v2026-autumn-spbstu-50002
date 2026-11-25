package currency

import (
	"encoding/xml"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
)

type XMLCurrency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type XMLData struct {
	Items []XMLCurrency `xml:"Valute"`
}

func ParseXML(filePath string) (*XMLData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open XML file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var data XMLData
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("could not decode XML data: %w", err)
	}

	return &data, nil
}
