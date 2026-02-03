package stringsx

/*
Реализовать пакет stringsx:
- нормализация строк
- безопасный split/join
- парсинг “key=value;key2=value2”
*/

// Normalize нормализует строку.
func Normalize(s string) string {
	var input = []rune(s)
	var output = make([]rune, 0)
	var start = 0

	for i := start; i < len(input) && input[i] == ' '; i++ {
		start++
	}

	for i := start; i < len(input); i++ {
		if input[i] != ' ' {
			output = append(output, input[i])
		}

		if input[i] == ' ' && input[i-1] != ' ' {
			output = append(output, ' ')
		}
	}

	if len(output) > 0 && output[len(output)-1] == ' ' {
		output = output[:len(output)-1]
	}

	return string(output)
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
