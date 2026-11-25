package currency

import (
	"fmt"
	"strconv"
	"strings"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ConvertData(source []XMLCurrency) ([]Currency, error) {
	result := make([]Currency, 0, len(source))

	for _, item := range source {
		valueStr := strings.ReplaceAll(item.Value, ",", ".")

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value format for currency %s: %w", item.CharCode, err)
		}

		result = append(result, Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	return result, nil
}
