package conversion_options

import (
	"fmt"
)

const Base = "Number Base"
const Distance = "Distance"

const Hexadecimal = "Hexadecimal"
const Decimal = "Decimal"
const Binary = "Binary"
const Roman = "Roman Numeral"

const Miles = "Miles"
const Kilometers = "Kilometers"
const Meters = "Meters"
const Yards = "Yards"

func GetTypesByCategory(category string) ([]string, error) {
	categoryMap := map[string][]string{
		Base: []string{Hexadecimal, Decimal, Binary, Roman},
		Distance: []string{Miles, Kilometers, Meters, Yards},
	}

	result, ok := categoryMap[category]
	if !ok {
		return nil, fmt.Errorf("Category not found: %s", category)
	}

	return result, nil
}

func GetConversionOptions(startingType string) ([]string, error) {

	conversionMap := map[string][]string{
		Hexadecimal: []string{Decimal, Binary, Roman},
		Decimal: []string{Hexadecimal, Binary, Roman},
		Binary: []string{Hexadecimal, Decimal, Roman},
		Roman: []string{Hexadecimal, Decimal, Binary},

		Miles: []string{Kilometers, Meters, Yards},
		Kilometers: []string{Miles, Meters, Yards},
		Meters: []string{Miles, Kilometers, Yards},
		Yards: []string{Miles, Kilometers, Meters},
	}

	result, ok := conversionMap[startingType]
	if !ok {
		return nil, fmt.Errorf("Unit or Base Type not Found: %s", startingType) 
	}

	return result, nil
}
