package stringsx

import "strings"

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
	return []string{}
}

// Join объединяет подстроки в строку с разделителем.
func Join(s []string, sep string) string {
	return ""
}

// ParseKV парсит строку в map.
func ParseKV(s string) map[string]string {
	return map[string]string{}
}
