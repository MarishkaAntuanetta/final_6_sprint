package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert — принимает строку и определяет, это код Морзе или обычный текст.
// Если код Морзе → переводит в текст.
// Если текст → переводит в код Морзе.
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("входная строка пустая")
	}

	if looksLikeMorse(input) {
		// Переводим код Морзе в текст
		return morse.ToText(input), nil
	} else {
		// Переводим текст в код Морзе
		return morse.ToMorse(input), nil
	}
}

// looksLikeMorse — проверяет, состоит ли строка только из точек, тире и пробелов
func looksLikeMorse(s string) bool {
	for _, ch := range s {
		if ch != '.' && ch != '-' && !unicode.IsSpace(ch) {
			return false
		}
	}
	return true
}
