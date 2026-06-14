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
// Маппинг позволяет сгруппировать несколько уровней в один (например, WARNING и WARN → WARN).
// По умолчанию нормализуются только WARNING → WARN.
// Для обработки любых уровней (CRITICAL, EMERGENCY, TRACE и т.д.) добавь их в mappings.
func NewLevelMapper() *LevelMapper {
	return &LevelMapper{
		mappings: map[string]string{
			"WARNING": "WARN",
		},
	}
}

// Normalize нормализует уровень логирования:
// - WARNING → WARN
// - Все остальные уровни возвращаются как есть (в верхнем регистре).
// Это позволяет обрабатывать любые уровни, встречающиеся в логах.
func (m *LevelMapper) Normalize(level string) string {
	upper := strings.ToUpper(level)
	if normalized, ok := m.mappings[upper]; ok {
		return normalized
	}
	return upper
}
