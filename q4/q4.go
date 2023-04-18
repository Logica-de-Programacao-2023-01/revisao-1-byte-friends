package q4

import "fmt"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	if taxCode >= 4 {
		return 0, fmt.Errorf("imposto náo encontrado")
	}
	if basePrice <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}
	var final float64
	switch taxCode {

	case 1:
		switch state {
		case "SP":
			final = basePrice + basePrice*0.05 + basePrice*0.1
			break
		case "RJ":
			final = basePrice + basePrice*0.05 + basePrice*0.15
			break
		case "MG":
			final = basePrice + basePrice*0.05 + basePrice*0.2
			break
		case "ES":
			final = basePrice + basePrice*0.05 + basePrice*0.25
			break
		case "OUTROS":
			final = basePrice + basePrice*0.05 + basePrice*0.3
		}

	case 2:
		switch state {
		case "SP":
			final = basePrice + basePrice*0.1 + basePrice*0.1
			break
		case "RJ":
			final = basePrice + basePrice*0.1 + basePrice*0.15
			break
		case "MG":
			final = basePrice + basePrice*0.1 + basePrice*0.2
			break
		case "ES":
			final = basePrice + basePrice*0.1 + basePrice*0.25
			break
		case "OUTROS":
			final = basePrice + basePrice*0.1 + basePrice*0.3
		}

	case 3:
		switch state {
		case "SP":
			final = basePrice + basePrice*0.15 + basePrice*0.1
			break
		case "RJ":
			final = basePrice + basePrice*0.15 + basePrice*0.15
			break
		case "MG":
			final = basePrice + basePrice*0.15 + basePrice*0.2
			break
		case "ES":
			final = basePrice + basePrice*0.15 + basePrice*0.25
			break
		case "OUTROS":
			final = basePrice + basePrice*0.15 + basePrice*0.3
		}
	}
	return final, nil
}
