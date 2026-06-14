// Package parser предоставляет интерфейсы и типы для парсинга логов.
// Поддерживает автоматическое определение формата и нормализацию уровней логирования.
package parser

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
