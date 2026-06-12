package parser

import (
	"errors"
	"regexp"
	"strings"

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
	// Используем (?i) для регистронезависимого поиска уровней
	// Проверяем, что после компонента есть ':'
	// WARN или WARNING, регистронезависимо
	pattern := `^(?i)(DEBUG|INFO|WARN(?:ING)?|WARNING|ERROR|FATAL|CRITICAL|EMERGENCY)\s+([a-zA-Z0-9_-]+):(.+)$`
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
		return nil, errors.New("empty line")
	}

	matches := p.re.FindStringSubmatch(trimmed)
	if matches == nil {
		return nil, errors.New("failed to parse custom format")
	}

	// matches[1] = level, matches[2] = component
	level := strings.ToUpper(matches[1]) // приводим к верхнему регистру для валидации
	component := matches[2]

	// Валидация уровня (проверяем, что это поддерживаемый уровень)
	if !IsSupportedLevel(level) {
		return nil, errors.New("unsupported log level")
	}

	// Используем LevelMapper для нормализации уровня
	normalizedLevel := p.levelMapper.Normalize(level)

	return &LogEntry{
		Level:     normalizedLevel,
		Component: component,
	}, nil
}
