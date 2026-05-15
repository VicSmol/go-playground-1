#
# Makefile for go-playground-1
#
# Цели:
#   make build     - Собрать бинарник
#   make run       - Собрать и запустить приложение
#   make test      - Запустить тесты с race detector и coverage
#   make vet       - Проверить код на ошибки статическим анализом
#   make fmt       - Отформатировать код
#   make lint      - Запустить линтер (если установлен golangci-lint)
#   make cover     - Показать отчёт по покрытию тестами
#   make doc       - Показать документацию по пакетам
#   make clean     - Удалить артефакты сборки
#   make all       - Выполнить все проверки (для CI)
#

# === Пути и переменные ===
BINARY := bin/app
GOFMT_FILES := $(shell find . -name "*.go" -not -path "*/.*" -not -path "*/vendor/*" -not -path "*/mocks/*")

# === Основные цели ===

.PHONY: all build run test vet fmt lint cover doc clean

all: fmt vet test cover
	@echo "✅ Все проверки пройдены: форматирование, веттинг, тесты, покрытие"

build: $(BINARY)

$(BINARY): $(GOFMT_FILES)
	go build -o $(BINARY) ./cmd

run: build
	./$(BINARY)

# test: запуск тестов с race detector и генерацией coverage
# как в CI
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# vet: статический анализ кода
vet:
	go vet ./...

# fmt: форматирование кода
fmt:
	@echo "Форматируем код..."
	gofmt -s -w $(GOFMT_FILES)

# lint: запуск golangci-lint (если установлен)
lint:
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint не установлен. Установите: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2"; \
		exit 1; \
	fi

# cover: показать отчёт по покрытию
cover:
	@if [ -f coverage.txt ]; then \
		go tool cover -func=coverage.txt; \
		echo ""; \
		go tool cover -html=coverage.txt -o coverage.html; \
		echo "➡️  Полный отчёт: coverage.html"; \
	else \
		echo "❌ Файл coverage.txt не найден. Сначала запустите 'make test'"; \
		exit 1; \
	fi

# doc: показать документацию по пакетам
doc:
	@echo "=== Документация по cmd ==="
	go doc ./cmd
	@echo "\n=== Документация по internal и всем подпакетам ==="
	@go list ./internal/... | xargs -r -n1 go doc

# clean: удалить артефакты сборки
clean:
	rm -rf $(BINARY) coverage.txt coverage.html

# === CI ===
# Используется в .github/workflows/ci.yml
# make ci = fmt + vet + test + cover
.PHONY: ci
ci: all
	@echo "✅ CI-проверки завершены успешно"
