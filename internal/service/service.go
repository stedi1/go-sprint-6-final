package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func CheckAndConvert(text string) string {
	// проверяем на морзе удалив все точки, запятые, пробелы
	str := strings.Trim(text, ".- ")

	if str == "" {
		// получили Азбуку морзе
		return morse.ToText(text)
	} else {
		// получили текст
		return morse.ToMorse(text)
	}
}
