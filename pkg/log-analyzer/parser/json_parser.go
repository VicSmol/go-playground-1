package parser

import (
	"encoding/json"
	"errors"
	"strings"
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
type JSONParser struct{}

// NewJSONParser создает новый экземпляр JSONParser.
func NewJSONParser() *JSONParser {
	return &JSONParser{}
}

// Parse разбирает JSON строку лога.
// Возвращает LogEntry или ошибку при невалидном JSON или отсутствии обязательных полей.
func (p *JSONParser) Parse(line string) (*LogEntry, error) {
	if strings.TrimSpace(line) == "" {
		return nil, errors.New("empty line")
	}

	// Проверка, что строка начинается с '{' и заканчивается '}'
	trimmed := strings.TrimSpace(line)
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

	return &LogEntry{
		Level:     NormalizeLevel(levelStr),
		Component: componentStr,
	}, nil
}
