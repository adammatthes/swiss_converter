package conversion_options

import (
	"fmt"
)

const Base = "Number Base"
const Distance = "Distance"
const Currency = "Currency"
const Temperature = "Temperature"
const Custom = "Custom"

const Hexadecimal = "Hexadecimal"
const Decimal = "Decimal"
const Binary = "Binary"
const Roman = "Roman Numeral"
const HexDec = Hexadecimal + Decimal
const HexBin = Hexadecimal + Binary
const HexRom = Hexadecimal + Roman
const DecHex = Decimal + Hexadecimal
const DecBin = Decimal + Binary
const DecRom = Decimal + Roman
const BinHex = Binary + Hexadecimal
const BinDec = Binary + Decimal
const BinRom = Binary + Roman
const RomHex = Roman + Hexadecimal
const RomDec = Roman + Decimal
const RomBin = Roman + Binary

const Miles = "Miles"
const Kilometers = "Kilometers"
const Meters = "Meters"
const Yards = "Yards"
const MilesKM = Miles + Kilometers
const MilesMeters = Miles + Meters
const MilesYards = Miles + Yards
const KMMiles = Kilometers + Miles
const KMMeters = Kilometers + Meters
const KMYards = Kilometers + Yards
const MetersMiles = Meters + Miles
const MetersKM = Meters + Kilometers
const MetersYards = Meters + Yards
const YardsKM = Yards + Kilometers
const YardsMiles = Yards + Miles
const YardsMeters = Yards + Meters

const USDollar = "usd"
const CanDollar = "cad"
const Peso = "mxn"
const Euro = "eur"
const USCan = USDollar + "-" + CanDollar
const USPeso = USDollar + "-" + Peso
const USEuro = USDollar + "-" + Euro
const CanUS = CanDollar + "-" + USDollar
const CanPeso = CanDollar + "-" + Peso
const CanEuro = CanDollar + "-" + Euro
const PesoUs = Peso + "-" + USDollar
const PesoCan = Peso + "-" + CanDollar
const PesoEuro = Peso + "-" + Euro
const EuroUS = Euro + "-" + USDollar
const EuroCan = Euro + "-" + CanDollar
const EuroPeso = Euro + "-" + Peso

const Fahrenheit = "Fahrenheit"
const Celsius = "Celsius"
const Kelvin = "Kelvin"
const FahCel = Fahrenheit + Celsius
const FahKel = Fahrenheit + Kelvin
const CelFah = Celsius + Fahrenheit
const CelKel = Celsius + Kelvin
const KelFah = Kelvin + Fahrenheit
const KelCel = Kelvin + Celsius

func GetAllTypes() []string {
	return []string{Hexadecimal, Decimal, Binary, Roman, Miles, Kilometers, Meters, Yards, USDollar, CanDollar, Peso, Euro, Fahrenheit, Celsius, Kelvin}
}

func GetTypesByCategory(category string) ([]string, error) {
	categoryMap := map[string][]string{
		Base: []string{Hexadecimal, Decimal, Binary, Roman},
		Distance: []string{Miles, Kilometers, Meters, Yards},
		Currency: []string{USDollar, CanDollar, Peso, Euro},
		Temperature: []string{Fahrenheit, Celsius, Kelvin},
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

		USDollar: []string{CanDollar, Peso, Euro},
		CanDollar: []string{USDollar, Peso, Euro},
		Peso: []string{USDollar, CanDollar, Euro},
		Euro: []string{USDollar, CanDollar, Peso},

		Fahrenheit: []string{Celsius, Kelvin},
		Celsius: []string{Fahrenheit, Kelvin},
		Kelvin: []string{Fahrenheit, Celsius},
	}

	result, ok := conversionMap[startingType]
	if !ok {
		return nil, fmt.Errorf("Unit or Base Type not Found: %s", startingType) 
	}

	return result, nil
}
