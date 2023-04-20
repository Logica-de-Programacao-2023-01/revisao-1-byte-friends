package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	if len(text) == 0 {
		return 0, fmt.Errorf("texto vazio")
	}
	caracteresEspeciais := []string{",", "!", ".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, caracter := range caracteresEspeciais {
		text = strings.ReplaceAll(text, caracter, "")
	}
	//
	var words []string
	words = strings.Fields(text)
	if len(words) == 0 {
		return 0, fmt.Errorf("texto vazio")
	}
	//
	var soma int = 0
	for _, letter := range words {
		soma += len(letter)
	}

	return float64(soma) / float64(len(words)), nil
}
