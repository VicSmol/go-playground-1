package parser

import (
	"strings"
)

// Dispatcher автоматически определяет формат лога и использует соответствующий парсер.
type Dispatcher struct {
	jsonParser    *JSONParser
	syslogParser  *SyslogParser
	customParser  *CustomParser
}

// NewDispatcher создает новый экземпляр Dispatcher.
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		jsonParser:    NewJSONParser(),
		syslogParser:  NewSyslogParser(),
		customParser:  NewCustomParser(),
	}
}

// Parse разбирает строку лога, автоматически определяя её формат.
// Поддерживаемые форматы:
// - JSON: строка начинается с '{'
// - Syslog: строка начинается с месяца (Jan|Feb|...|Dec)
// - Custom: строка начинается с уровня (DEBUG|INFO|WARN|WARNING|ERROR|FATAL|CRITICAL|EMERGENCY)
func (d *Dispatcher) Parse(line string) (*LogEntry, error) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return nil, nil // Пропускаем пустые строки
	}

	// Автоматическое определение формата

	// 1. Проверка на JSON
	if strings.HasPrefix(trimmed, "{") {
		return d.jsonParser.Parse(line)
	}

	// 2. Проверка на Syslog (начинается с месяца)
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	for _, month := range months {
		if strings.HasPrefix(trimmed, month) {
			return d.syslogParser.Parse(line)
		}
	}

	// 3. Проверка на Custom формат (начинается с уровня)
	levels := []string{"DEBUG", "INFO", "WARN", "WARNING", "ERROR", "FATAL", "CRITICAL", "EMERGENCY"}
	for _, level := range levels {
		if strings.HasPrefix(trimmed, level+" ") || strings.HasPrefix(trimmed, level+":") {
			return d.customParser.Parse(line)
		}
	}

	// Не удалось определить формат
	return nil, nil
}

// ParseWithParser разбирает строку с использованием указанного парсера.
// Используется для тестирования конкретного парсера.
func (d *Dispatcher) ParseWithParser(parser LogParser, line string) (*LogEntry, error) {
	return parser.Parse(line)
}
