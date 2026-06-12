// Package mappers предоставляет мапперы для преобразования данных из логов.
package mappers

import (
	"strings"
)

// LevelMapper маппер для нормализации уровней логирования.
type LevelMapper struct {
	mappings map[string]string
}

// NewLevelMapper создает новый экземпляр LevelMapper с инициализированным маппингом.
func NewLevelMapper() *LevelMapper {
	return &LevelMapper{
		mappings: map[string]string{
			"DEBUG":     "DEBUG",
			"INFO":      "INFO",
			"WARN":      "WARN",
			"WARNING":   "WARN",
			"ERROR":     "ERROR",
			"FATAL":     "FATAL",
			"CRITICAL":  "FATAL",
			"EMERGENCY": "FATAL",
		},
	}
}

// Normalize нормализует уровень логирования:
// - WARNING → WARN
// - CRITICAL, FATAL, EMERGENCY → FATAL
// Остальные уровни возвращаются как есть (в верхнем регистре).
func (m *LevelMapper) Normalize(level string) string {
	upper := strings.ToUpper(level)
	if normalized, ok := m.mappings[upper]; ok {
		return normalized
	}
	return upper
}
