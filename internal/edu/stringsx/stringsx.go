package stringsx

import (
	"strings"
)

/*
Реализовать пакет stringsx:
- нормализация строк
- безопасный split/join
- парсинг “key=value;key2=value2”
*/

// Normalize нормализует строку.
func Normalize(s string) string {
	var input = []rune(s)
	var builder strings.Builder
	var isLastSpace bool

	for _, r := range input {
		if !isLastSpace && r == ' ' {
			builder.WriteRune(r)
			isLastSpace = true
		}

		if r != ' ' {
			builder.WriteRune(r)
			isLastSpace = false
		}
	}

	return strings.TrimSpace(builder.String())
}

// Split разбивает строку на подстроки по разделителю.
func Split(s string, sep string) []string {
	separator := []rune(sep)[0]
	symbols := []rune(s)
	length := len(symbols)
	result := make([]string, 0)
	var builder strings.Builder

	for i := 0; i < length; i++ {
		if symbols[i] == separator && builder.Len() > 0 {
			result = append(result, builder.String())
			builder.Reset()

			continue
		}

		if symbols[i] != separator {
			builder.WriteRune(symbols[i])
		}
	}

	if builder.Len() > 0 {
		result = append(result, builder.String())
	}

	return result
}

// Join объединяет подстроки в строку с разделителем.
func Join(s []string, sep string) string {
	var separator = []rune(sep)[0]
	var builder strings.Builder

	if len(s) == 0 {
		return ""
	}

	for index, str := range s {
		builder.WriteString(str)

		if index != len(s)-1 {
			builder.WriteRune(separator)
		}
	}

	return builder.String()
}

// ParseKV парсит строку в map.
func ParseKV(s string) map[string]string {
	separator := ";"
	splits := Split(s, separator)
	result := map[string]string{}

	for _, split := range splits {
		keyValue := Split(split, "=")
		if len(keyValue) == 2 {
			result[keyValue[0]] = keyValue[1]
		}
	}

	return result
}
