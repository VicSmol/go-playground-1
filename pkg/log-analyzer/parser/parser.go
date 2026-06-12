// Package parser предоставляет интерфейсы и типы для парсинга логов.
// Поддерживает автоматическое определение формата и нормализацию уровней логирования.
package parser

import (
	"strings"
)

// LogEntry представляет собой одну запись лога после парсинга.
type LogEntry struct {
	Level     string // Нормализованный уровень логирования (DEBUG, INFO, WARN, ERROR, FATAL)
	Component string // Идентификатор компонента, сгенерировавшего лог
}

// LogParser определяет интерфейс для парсинга логов.
type LogParser interface {
	// Parse разбирает одну строку лога и возвращает LogEntry или ошибку.
	// При ошибке парсинга (некорректный формат) возвращает (nil, error).
	Parse(line string) (*LogEntry, error)
}

// NormalizeLevel нормализует уровень логирования согласно ТЗ:
// - WARNING → WARN
// - CRITICAL, FATAL, EMERGENCY → FATAL
// Остальные уровни возвращаются как есть (в верхнем регистре).
func NormalizeLevel(level string) string {
	upper := strings.ToUpper(level)
	switch upper {
	case "WARNING":
		return "WARN"
	case "CRITICAL", "FATAL", "EMERGENCY":
		return "FATAL"
	default:
		return upper
	}
}

// SupportedLevels возвращает список поддерживаемых уровней логирования.
func SupportedLevels() []string {
	return []string{"DEBUG", "INFO", "WARN", "WARNING", "ERROR", "FATAL", "CRITICAL", "EMERGENCY"}
}

// IsSupportedLevel проверяет, является ли уровень логирования поддерживаемым.
func IsSupportedLevel(level string) bool {
	upper := strings.ToUpper(level)
	for _, supported := range SupportedLevels() {
		if upper == supported {
			return true
		}
	}
	return false
}
