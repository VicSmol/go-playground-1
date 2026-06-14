package parser

import (
	"regexp"
	"strings"

	"go-playground-1/internal/log-analyzer/errors"
	"go-playground-1/internal/log-analyzer/mappers"
)

// CustomParser парсит логи в кастомном формате: LEVEL COMPONENT: message
// Пример: INFO auth: User logged in
type CustomParser struct {
	// Регулярное выражение для парсинга custom формата
	re *regexp.Regexp
	// LevelMapper для нормализации уровней логирования
	levelMapper *mappers.LevelMapper
}

// NewCustomParser создает новый экземпляр CustomParser.
func NewCustomParser() *CustomParser {
	// Pattern: LEVEL COMPONENT: message
	// Уровень в начале, component до первого ':'
	// Уровень может быть любым словом, состоящим из букв, цифр, подчеркиваний или дефисов
	pattern := `^([A-Za-z0-9_-]+)\s+([a-zA-Z0-9_-]+):(.+)$`
	return &CustomParser{
		re:          regexp.MustCompile(pattern),
		levelMapper: mappers.NewLevelMapper(),
	}
}

// Parse разбирает custom формат лога.
// Возвращает LogEntry или ошибку при невалидном формате.
func (p *CustomParser) Parse(line string) (*LogEntry, error) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return nil, errors.ParserErrorEmptyLine
	}

	matches := p.re.FindStringSubmatch(trimmed)
	if matches == nil {
		return nil, errors.ParserErrorInvalidCustomFormat
	}

	// matches[1] = level, matches[2] = component
	level := strings.ToUpper(matches[1]) // приводим к верхнему регистру для валидации
	component := matches[2]

	// Используем LevelMapper для нормализации уровня
	normalizedLevel := p.levelMapper.Normalize(level)

	return &LogEntry{
		Level:     normalizedLevel,
		Component: component,
	}, nil
}
