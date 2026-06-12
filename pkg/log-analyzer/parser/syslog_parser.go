package parser

import (
	"errors"
	"regexp"
	"strings"
)

// SyslogParser парсит логи в формате RFC 3164.
// Формат: Mon DD HH:MM:SS host component[pid]: [LEVEL] message
// Пример: Jun  8 10:00:00 host nginx[123]: [INFO] User logged in
type SyslogParser struct {
	// Регулярное выражение для парсинга syslog
	re *regexp.Regexp
}

// NewSyslogParser создает новый экземпляр SyslogParser.
func NewSyslogParser() *SyslogParser {
	// Pattern: Month Day Time Hostname Component[PID]: [LEVEL] Message
	// Пример: Jun  8 10:00:00 host nginx[123]: [INFO] User logged in
	pattern := `^([A-Z][a-z]{2})\s+(\d{1,2})\s+(\d{2}:\d{2}:\d{2})\s+(\S+)\s+(\S+?):?\s*\[?([A-Z]+)\]?`
	return &SyslogParser{
		re: regexp.MustCompile(pattern),
	}
}

// Parse разбирает syslog строку.
// Возвращает LogEntry или ошибку при невалидном формате.
func (p *SyslogParser) Parse(line string) (*LogEntry, error) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return nil, errors.New("empty line")
	}

	// Проверка, что строка начинается с месяца
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	startsWithMonth := false
	for _, month := range months {
		if strings.HasPrefix(trimmed, month) {
			startsWithMonth = true
			break
		}
	}
	if !startsWithMonth {
		return nil, errors.New("not a syslog line")
	}

	matches := p.re.FindStringSubmatch(trimmed)
	if matches == nil {
		return nil, errors.New("failed to parse syslog format")
	}

	// matches[5] = component (может содержать [pid]), matches[6] = level
	component := matches[5]
	level := matches[6]

	// Удаляем PID из компонента, если он есть (например, nginx[123] -> nginx)
	if idx := strings.Index(component, "["); idx != -1 {
		component = component[:idx]
	}

	// Валидация уровня
	if !IsSupportedLevel(level) {
		return nil, errors.New("unsupported log level")
	}

	return &LogEntry{
		Level:     NormalizeLevel(level),
		Component: component,
	}, nil
}
