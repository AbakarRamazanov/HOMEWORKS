package hw02unpackstring

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(text string) (string, error) {
	mas := []rune(text)
	var textReturn string
	var number int

	for i, symbol := range mas {
		if unicode.IsDigit(symbol) && i == 0 {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(symbol) && unicode.IsDigit(mas[i-1]) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(symbol) {
			number = int(symbol - '0')
			if number == 0 {
				// textReturn = textReturn[:utf8.RuneCountInString(textReturn)-1]
				continue
			}
			for j := 0; j < number-1; j++ {
				textReturn += string(mas[i-1])
			}
			continue
		}
		textReturn += string(symbol)
	}
	return textReturn, nil
}
