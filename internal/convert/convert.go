package convert

import (
	"strconv"
	"fmt"
	"github.com/adammatthes/swiss_converter/internal/conversion_options"
)

func HexadecimalToDecimal(input string) (string, error) {
	if input[0] == '-' {
		input = input[1:]
	}

	i64, err := strconv.ParseUint(input, 16, 64)
	if err != nil {
		return "", fmt.Errorf("Invalid Number for conversion: %s", input)
	}

	return fmt.Sprintf("%v", i64), nil
}

func GetConversionFunction(start, end string) (func(string) (string, error), error) {
	functions := map[string]func(string) (string, error) {
		conversion_options.Hexadecimal + conversion_options.Decimal: HexadecimalToDecimal,
	}

	result, ok := functions[start+end]
	if !ok {
		return nil, fmt.Errorf("Conversion function not found")
	}

	return result, nil
}
