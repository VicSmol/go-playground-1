package parser

import (
	"errors"
	"regexp"
	"strings"
)

// CustomParser парсит логи в кастомном формате: LEVEL COMPONENT: message
// Пример: INFO auth: User logged in
type CustomParser struct {
	// Регулярное выражение для парсинга custom формата
	re *regexp.Regexp
}

// NewCustomParser создает новый экземпляр CustomParser.
func NewCustomParser() *CustomParser {
	// Pattern: LEVEL COMPONENT: message
	// Уровень в начале, component до первого ':'
	// Используем (?i) для регистронезависимого поиска уровней
	// Проверяем, что после компонента есть ':'
	// WARN или WARNING, регистронезависимо
	pattern := `^(?i)(DEBUG|INFO|WARN(?:ING)?|WARNING|ERROR|FATAL|CRITICAL|EMERGENCY)\s+([a-zA-Z0-9_-]+):(.+)$`
	return &CustomParser{
		re: regexp.MustCompile(pattern),
	}
}

// Parse разбирает custom формат лога.
// Возвращает LogEntry или ошибку при невалидном формате.
func (p *CustomParser) Parse(line string) (*LogEntry, error) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return nil, errors.New("empty line")
	}

	matches := p.re.FindStringSubmatch(trimmed)
	if matches == nil {
		return nil, errors.New("failed to parse custom format")
	}

	// matches[1] = level, matches[2] = component
	level := matches[1]
	component := matches[2]

	// Валидация уровня (проверяем, что это поддерживаемый уровень)
	if !IsSupportedLevel(level) {
		return nil, errors.New("unsupported log level")
	}

	return &LogEntry{
		Level:     NormalizeLevel(level),
		Component: component,
	}, nil
}
