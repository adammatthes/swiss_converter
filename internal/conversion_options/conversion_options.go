package conversion_options

import (
	"fmt"
)

func GetConversionOptions(startingType string) ([]string, error) {
	hexadecimal := "Hexadecimal"
	decimal := "Decimal"
	binary := "Binary"
	roman := "Roman Numeral"

	conversionMap := map[string][]string{
		hexadecimal: []string{decimal, binary, roman},
		decimal: []string{hexadecimal, binary, roman},
		binary: []string{hexadecimal, decimal, roman},
		roman: []string{hexadecimal, decimal, binary},
	}

	result, ok := conversionMap[startingType]
	if !ok {
		return nil, fmt.Errorf("Unit or Base Type not Found: %s", startingType) 
	}

	return result, nil
}
