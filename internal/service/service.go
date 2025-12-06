package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func CheckAndConvert(text string) (string, error) {
	// проверка на отсутсвие текста
	if len(text) < 1 {
		return "", errors.New("no any text")
	}
	// запрещаем латинские буквы
	eng := strings.ContainsAny(strings.ToLower(text), "abcdefghijklmnopqrstuvwxyz")
	if eng {
		return "", errors.New("latin characters are prohibited")
	}

	// проверяем на морзе удалив все точки, запятые, пробелы
	str := strings.Trim(text, ".- ")

	if str == "" {
		// получили Азбуку морзе
		return morse.ToText(text), nil
	} else {
		// получили текст
		return morse.ToMorse(text), nil
	}
}
