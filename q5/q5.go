package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	if fromScale == "" && toScale == "" {
		return 0, fmt.Errorf("escala inválida")
	}
	//
	var result float64
	switch fromScale {
	case "C":
		if toScale == "F" {
			result = temp*1.8 + 32
		} else if toScale == "K" {
			result = temp + 273.15
		}
	case "F":
		if toScale == "C" {
			result = (temp - 32) * 0.55
		} else if toScale == "K" {
			result = (temp-32)*0.55 + 273.15
		}
	case "K":
		if toScale == "C" {
			result = temp - 273.15
		} else if toScale == "F" {
			result = (temp-273.15)*1.8 + 32
		}
	}

	return result, nil
}
