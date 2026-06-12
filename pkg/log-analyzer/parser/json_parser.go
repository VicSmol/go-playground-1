package parser

import (
	"encoding/json"
	"errors"
	"strings"

	"go-playground-1/internal/log-analyzer/mappers"
)

// ErrMissingLevel occurs when JSON log entry doesn't have a level field.
var ErrMissingLevel = errors.New("missing required field: level")

// ErrMissingComponent occurs when JSON log entry doesn't have a component field.
var ErrMissingComponent = errors.New("missing required field: component")

// ErrEmptyLevel occurs when level field is empty.
var ErrEmptyLevel = errors.New("level field is empty")

// ErrEmptyComponent occurs when component field is empty.
var ErrEmptyComponent = errors.New("component field is empty")

// JSONParser парсит логи в формате JSON.
// Ожидает JSON с полями: level (обязательное), component (обязательное), message, timestamp (опциональные).
type JSONParser struct {
	// LevelMapper для нормализации уровней логирования
	levelMapper *mappers.LevelMapper
}

// NewJSONParser создает новый экземпляр JSONParser.
func NewJSONParser() *JSONParser {
	return &JSONParser{
		levelMapper: mappers.NewLevelMapper(),
	}
}

// Parse разбирает JSON строку лога.
// Возвращает LogEntry или ошибку при невалидном JSON или отсутствии обязательных полей.
func (p *JSONParser) Parse(line string) (*LogEntry, error) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return nil, errors.New("empty line")
	}

	// Проверка, что строка начинается с '{' и заканчивается '}'
	if !strings.HasPrefix(trimmed, "{") {
		return nil, errors.New("not a JSON line")
	}
	if !strings.HasSuffix(trimmed, "}") {
		return nil, errors.New("invalid JSON: missing closing brace")
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(line), &data); err != nil {
		return nil, err
	}

	// Проверка обязательных полей
	level, ok := data["level"]
	if !ok {
		return nil, ErrMissingLevel
	}
	levelStr, ok := level.(string)
	if !ok || levelStr == "" {
		return nil, ErrEmptyLevel
	}

	component, ok := data["component"]
	if !ok {
		return nil, ErrMissingComponent
	}
	componentStr, ok := component.(string)
	if !ok || componentStr == "" {
		return nil, ErrEmptyComponent
	}

	// Используем LevelMapper для нормализации уровня
	normalizedLevel := p.levelMapper.Normalize(levelStr)

	return &LogEntry{
		Level:     normalizedLevel,
		Component: componentStr,
	}, nil
}
