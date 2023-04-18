package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	if len(text) == 0 {
		return 0, fmt.Errorf("texto vazio")
	}
	// limpagem
	slice := []string{""}
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "!", "")

	// separar as palavras e coloca em uma slice
	slice = strings.Fields(text)
	var words float64
	words = float64(len(slice))

	//quantidade de letras
	var letras float64
	text = strings.ReplaceAll(text, " ", "")
	letras = float64(len(text)) - 2 //subtrair acentos

	var media float64
	media = letras / words

	return media, nil
}
