package convert

import (
	"strconv"
	"fmt"
	"github.com/adammatthes/swiss_converter/internal/conversion_options"
)

func getFloat(input string) (float64, error) {
	if input[0] == '-' {
		input = input[1:]
	}

	f64, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0, err
	}

	return f64, nil
}

func getInt(input string, base int) (uint64, error) {
	if input[0] == '-' {
		input = input[1:]
	}

	i64, err := strconv.ParseUint(input, base, 64)
	if err != nil {
		return 0, err
	}

	return i64, nil
}

func HexadecimalToDecimal(input string) (string, error) {
	i64, err := getInt(input, 16)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", i64), nil
}

func HexadecimalToBinary(input string) (string, error) {
	i64, err := getInt(input, 16)
	if err != nil {
		return "", fmt.Errorf("Invalid Number for conversion: %s", input)
	}

	return fmt.Sprintf("%b", i64), nil
}

func DecimalToRoman(input string) (string, error) {
	i64, err := getInt(input, 10)
	if err != nil {
		return "", fmt.Errorf("Invalid Number for conversion: %s", input)
	}

	if i64 >= 4000 {
		return "", fmt.Errorf("Roman numeral conversion does not yet support numbers larger than 3999")
	}

	result := ""
	for ; i64 >= 1000; i64 -= 1000 {
		result += "M"
	}

	if i64 >= 900 {
		result += "CM"
		i64 -= 900
	}

	if i64 >= 500 {
		result += "D"
		i64 -= 500
	}

	for ; i64 >= 100; i64 -= 100 {
		result += "C"
	}

	if i64 >= 90 {
		result += "XC"
		i64 -= 90
	}

	if i64 >= 50 {
		result += "L"
		i64 -= 50
	}

	for ; i64 >= 10; i64 -= 10 {
		result += "X"
	}

	if i64 == 9 {
		result += "IX"
		i64 -= 9
	}

	if i64 >= 5 {
		result += "V"
		i64 -= 5
	}

	for ; i64 > 0; i64 -= 1 {
		result += "I"
	}

	return result, nil
}

func HexadecimalToRoman(input string) (string, error) {
	dec, err := HexadecimalToDecimal(input)
	if err != nil {
		return "", err
	}

	result, err := DecimalToRoman(dec)
	if err != nil {
		return "", err
	}

	return result, nil
}

func DecimalToHexadecimal(input string) (string, error) {
	i64, err := getInt(input, 10)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%X", i64)
	return result, nil
}

func DecimalToBinary(input string) (string, error) {
	i64, err := getInt(input, 10)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%b", i64)
	return result, nil
}

func BinaryToHexadecimal(input string) (string, error) {
	i64, err := getInt(input, 2)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%X", i64)
	return result, nil
}

func BinaryToDecimal(input string) (string, error) {
	i64, err := getInt(input, 2)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", i64)
	return result, nil
}

func BinaryToRoman(input string) (string, error) {
	dec, err := BinaryToDecimal(input)
	if err != nil {
		return "", err
	}

	result, err := DecimalToRoman(dec)
	if err != nil {
		return "", err
	}

	return result, nil
}

func MilesToKilometers(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 1.60934)
	return result, nil
}

func KilometersToMeters(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 1000.0)
	return result, nil
}

func MilesToMeters(input string) (string, error) {
	km, err := MilesToKilometers(input)
	if err != nil {
		return "", err
	}

	result, err := KilometersToMeters(km)
	if err != nil {
		return "", err
	}

	return result, nil
}

func MilesToYards(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 1760.0)
	return result, nil
}

func KilometersToMiles(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 0.621371)
	return result, nil
}

func KilometersToYards(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 1093.61)
	return result, nil
}

func YardsToKilometers(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 0.0009144)
	return result, nil
}

func YardsToMiles(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 0.000568182)
	return result, nil
}

func YardsToMeters(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 0.9144)
	return result, nil
}

func MetersToMiles(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 0.000621371)
	return result, nil
}

func MetersToKilometers(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", f64 * 0.001)
	return result, nil
}

func MetersToYards(input string) (string, error) {
	f64, err := getFloat(input)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("$v", f64 * 1.09361)
	return result, nil
}

func GetConversionFunction(start, end string) (func(string) (string, error), error) {
	functions := map[string]func(string) (string, error) {
		conversion_options.HexDec: HexadecimalToDecimal,
		conversion_options.HexBin: HexadecimalToBinary,
		conversion_options.DecRom: DecimalToRoman,
		conversion_options.HexRom: HexadecimalToRoman,
		conversion_options.DecHex: DecimalToHexadecimal,
		conversion_options.DecBin: DecimalToBinary,
		conversion_options.BinHex: BinaryToHexadecimal,
		conversion_options.BinDec: BinaryToDecimal,
		conversion_options.BinRom: BinaryToRoman,
		conversion_options.MilesKM: MilesToKilometers,
		conversion_options.MilesMeters: MilesToMeters,
		conversion_options.MilesYards: MilesToYards,
		conversion_options.KMMiles: KilometersToMiles,
		conversion_options.KMMeters: KilometersToMeters,
		conversion_options.KMYards: KilometersToYards,
		conversion_options.YardsKM: YardsToKilometers,
		conversion_options.YardsMiles: YardsToMiles,
		conversion_options.YardsMeters: YardsToMeters,
		conversion_options.MetersMiles: MetersToMiles,
		conversion_options.MetersKM: MetersToKilometers,
		conversion_options.MetersYards: MetersToYards,
	}

	result, ok := functions[start+end]
	if !ok {
		return nil, fmt.Errorf("Conversion function not found")
	}

	return result, nil
}
