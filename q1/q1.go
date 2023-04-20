package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	if currentPurchase <= 0 {
		return 0, fmt.Errorf("valor de compra inválido")
	}
	var soma float64
	soma = 0
	for _, sum := range purchaseHistory {
		soma = soma + sum
	}
	var media float64
	media = soma / float64(len(purchaseHistory))

	var desconto float64
	if len(purchaseHistory) == 0 {
		desconto = currentPurchase * 0.1
	} else if media > 1000 {
		desconto = currentPurchase * 0.2
	} else if soma > 1000 { //
		desconto = currentPurchase * 0.1
	} else if currentPurchase == purchaseHistory[0] {
		desconto = currentPurchase * 0.1
	} else if soma <= 500 {
		desconto = currentPurchase * 0.02
	} else if soma <= 1000 {
		desconto = currentPurchase * 0.05
	}
	return desconto, nil
}
