package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert — принимает строку и определяет, это код Морзе или обычный текст.
// Если код Морзе переводит в текст.
// Если текст переводит в код Морзе.
func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("the input string is empty")
	}

	if isMorse(input) {
		// Переводим код Морзе в текст
		return morse.ToText(input), nil
	} else {
		// Переводим текст в код Морзе
		return morse.ToMorse(input), nil
	}
}

// looksLikeMorse — проверяет, состоит ли строка только из точек, тире и пробелов
func isMorse(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	// Ищем первый символ, который НЕ является ., - или пробелом
	idx := strings.IndexFunc(s, func(r rune) bool {
		return r != '.' && r != '-' && !unicode.IsSpace(r)
	})

	// Если такого символа нет — это код Морзе
	return idx == -1
}
