package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func ParseXML(data []byte) (ValCurs, error) {
	var valCurs ValCurs
	if err := xml.Unmarshal(data, &valCurs); err != nil {
		return ValCurs{}, err
	}

	return valCurs, nil
}

func parseValue(value string) (float64, error) {
	normalized := strings.ReplaceAll(value, ",", ".")

	return strconv.ParseFloat(normalized, 64)
}

type CurrencyResult struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ToResults(valutes []Valute) ([]CurrencyResult, error) {
	results := make([]CurrencyResult, 0, len(valutes))
	for _, v := range valutes {
		numCode, err := strconv.Atoi(v.NumCode)
		if err != nil {
			return nil, fmt.Errorf("invalid NumCode %q: %w", v.NumCode, err)
		}

		value, err := parseValue(v.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid Value %q: %w", v.Value, err)
		}

		results = append(results, CurrencyResult{
			NumCode:  numCode,
			CharCode: v.CharCode,
			Value:    value,
		})
	}

	return results, nil
}
