# Spec-First Подход в Разработке ПО: Полное Руководство

Spec-First (также известный как Design-First или Contract-First) — это парадигма разработки программного обеспечения, при которой **спецификация (specification)** системы, её интерфейса или API создается **до написания кода**. Этот документ становится **единственным источником истины (single source of truth)**, от которого отталкиваются все участники проекта.

---


## 1. Суть и философия Spec-First

### 1.1. Принцип

Spec-First противопоставляется традиционному **Code-First** подходу:

*   **Code-First:** `Код → (Документация, Тесты)` 
*   **Spec-First:** `Спецификация → (Код, Тесты, Документация)`


Цель Spec-First — **сдвинуть согласование и проектирование на самый ранний этап**, минимизируя затраты на пересогласование и рефакторинг в дальнейшем.

### 1.2. Ключевые преимущества

| Преимущество | Описание |
|------------|---------|
| **Раннее согласование** | Фронтенд, бэкенд, QA, DevOps, бизнес-аналитики работают по единой, заранее согласованной спецификации. |
| **Снижение издержек** | Изменения в API обсуждаются на уровне spec, что дешевле и быстрее, чем в коде. |
| **Параллельная разработка** | Команды могут работать независимо: фронтенд использует сгенерированный мок, бэкенд — реализует заглушки. |
| **Гарантированная согласованность** | Вся документация и код происходят из одного источника, что исключает расхождения. |
| **Автоматизация** | CI/CD может проверять обратную совместимость спецификации. |

---


## 2. Spec-First для веб-API: OpenAPI (Swagger)

### 2.1. Основные форматы

*   **OpenAPI 3.0:** Стандарт де-факто для описания RESTful API. Позволяет описать:
    *   Эндпоинты (paths).
    *   Параметры запроса, заголовки, тела.
    *   Схемы запросов и ответов (через JSON Schema).
    *   Коды ошибок (404, 500 и т.д.).
    *   Автоматически генерируемую документацию.


**Пример спецификации (`api-spec.yaml`):**
```yaml
openapi: 3.0.3
info:
  title: Todo API
  version: 1.0.0
servers:
  - url: http://localhost:8080/v1

paths:
  /todos:
    get:
      summary: Получить список задач
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Todo'
    post:
      summary: Создать новую задачу
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/CreateTodoRequest'
      responses:
        '201':
          description: Задача создана
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Todo'

components:
  schemas:
    Todo:
      type: object
      properties:
        id:
          type: integer
        title:
          type: string
        completed:
          type: boolean
    CreateTodoRequest:
      type: object
      required:
        - title
      properties:
        title:
          type: string
        description:
          type: string
```

### 2.2. Инструменты и генерация кода

На основе `api-spec.yaml` можно автоматически сгенерировать:

*   **Серверный код (Go, Python, Java и т.д.):** Заглушки обработчиков (handlers).
*   **Клиентские SDK:** Для разных языков.
*   **Документацию (Swagger UI, Redoc):** Интерактивный веб-интерфейс.
*   **Тесты и Mock-серверы.**

**Популярные инструменты:**
*   [OpenAPI Generator](https://openapi-generator.tech/)
*   [Swagger Codegen](https://swagger.io/tools/swagger-codegen/)

---


## 3. Spec-First для консольных приложений (CLI)

Хотя для CLI нет единого стандарта, как OpenAPI, принципы Spec-First применяются аналогично.


### 3.1. Определение спецификации CLI

Спецификация CLI описывает команды, подкоманды, флаги, аргументы и формат вывода. Её можно представить в виде структурированного файла.


**Пример спецификации (`cli-spec.yaml`):**
```yaml
name: todo-cli
version: "1.0"
description: Утилита командной строки для управления задачами

commands:
  - name: add
    description: Добавить новую задачу
    args:
      - name: title
        description: Заголовок задачи
        required: true
    flags:
      - name: priority
        short: p
        type: string
        enum: [low, normal, high]
        default: "normal"
        description: Приоритет задачи
      - name: due
        type: string
        description: Срок выполнения (YYYY-MM-DD)

  - name: list
    description: Показать список задач
    flags:
      - name: status
        type: string
        enum: [todo, in-progress, done]
        description: Фильтр по статусу
      - name: priority
        short: p
        type: string
        enum: [low, normal, high]
        description: Фильтр по приоритету

  - name: complete
    description: Отметить задачу как выполненную
    args:
      - name: id
        description: ID задачи
        required: true

  - name: help
    description: Показать справку
```

### 3.2. Инструменты и генерация для CLI

#### a) **Генерация кода и документации**

На основе `cli-spec.yaml` можно написать скрипт, который генерирует:

*   **Каркас приложения на Go (с использованием `cobra`):** Создает файлы для каждой команды.
*   **Документацию (`README.md`):** Описание всех команд и флагов.
*   **Автодополнение (bash completion).**


**Пример генерации команды `add` (Go):**
```go
// cmd/add.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var priority string
var due string

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Добавить новую задачу",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]
		// ... логика добавления задачи
		fmt.Printf("Задача '%s' добавлена с приоритетом %s\n", title, priority)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&priority, "priority", "p", "normal", "Приоритет задачи (low, normal, high)")
	addCmd.MarkFlagRequired("priority") // Но в spec он не required, это пример
	addCmd.Flags().StringVar(&due, "due", "", "Срок выполнения (YYYY-MM-DD)")
	rootCmd.AddCommand(addCmd)
}
```

#### b) **Популярные фреймворки**

*   **[Cobra](https://github.com/spf13/cobra) (Go):** Позволяет декларативно описывать команды. Хотя он не генерирует код из YAML, его структура отражает проект, как если бы это была спецификация.
*   **[oclif](https://oclif.io/) (Node.js):** Поддерживает генерацию CLI-приложений из конфигурации.
*   **[clikt](https://ajalt.github.io/clikt/) (Kotlin):** Позволяет описывать CLI с помощью аннотаций.


### 3.3. Преимущества для CLI

*   **Предсказуемый UX:** Все команды и флаги согласованы по стилю.
*   **Актуальная справка:** `--help` всегда соответствует спецификации.
*   **Легкость в разработке:** Разработчики реализуют только бизнес-логику.
*   **Простота для пользователей:** Четкая, автоматически сгенерированная документация.

---


## 4. Лучшие практики и рекомендации

1.  **Интегрируйте в CI/CD:** Проверяйте, что изменения в спецификации не нарушают обратную совместимость.
2.  **Используйте версионирование:** Управляйте изменениями в спецификации (например, через Git).
3.  **Документируйте политику изменений:** Какие изменения считаются обратно совместимыми?
4.  **Выбирайте подходящий инструмент:** Для веб-API — OpenAPI; для CLI — YAML + генератор или фреймворк вроде Cobra.
5.  **Начинайте с малого:** Не нужно описывать всё сразу. Начните с ключевых эндпоинтов или команд.

---


## 5. Сравнение: Spec-First vs Code-First

| Характеристика | **Spec-First** | **Code-First** |
|---------------|----------------|----------------|
| **Стартовая точка** | Спецификация | Код |
| **Гибкость** | Ниже (жёсткая договорённость) | Выше (быстрое прототипирование) |
| **Масштабируемость** | Отличная (для больших команд) | Сложнее |
| **Документация** | Автоматически актуальна | Часто устаревает |
| **Идеальный сценарий** | Крупные проекты, публичные API, микросервисы | PoC, скрипты, внутренние инструменты |

---


## 6. Spec-First для других компонентов

### 6.1. Контракты с Message Brokers (AsyncAPI)

В мире микросервисов и event-driven архитектур общение часто происходит через message brokers (Kafka, RabbitMQ, NATS). Для описания таких контрактов используется **AsyncAPI**.

**AsyncAPI** — это аналог OpenAPI для асинхронных API. Он позволяет описать:

*   **Брокер сообщений:** Тип (Kafka, RabbitMQ), хост, порт, безопасность.
*   **Каналы (Channels):** Аналог эндпоинтов в REST. Например, `user/signedup`.
*   **Операции:** `publish` (отправка сообщения) и `subscribe` (подписка на сообщения).
*   **Схемы сообщений:** Формат данных, публикуемых или подписываемых на канале.

*   **Привязки:** Специфичные для брокера настройки (например, `kafka: { groupId: user-service }`).


**Пример спецификации (`asyncapi.yaml`):**
```yaml
asyncapi: 2.6.0
info:
  title: User Sign-Up Events
  version: 1.0.0
  description: События регистрации пользователей

servers:
  production:
    url: kafka-prod.example.com:9092
    protocol: kafka
    protocolVersion: "1.0.0"

channels:
  user/signedup:
    subscribe:
      message:
        $ref: '#/components/messages/UserSignedUp'

components:
  messages:
    UserSignedUp:
      payload:
        type: object
        properties:
          userId:
            type: string
          email:
            type: string
            format: email
          timestamp:
            type: string
            format: date-time
```

**Преимущества и использование:**
*   Генерация кода для producer и consumer.
*   Автоматическая генерация документации и DDL-скриптов (например, для создания топиков Kafka).
*   Мокирование брокера для тестирования.

---


### 6.2. Контракты с базой данных

Spec-First также применим к схемам баз данных, что часто называют **Schema-as-Code** или **Database-First**.

**Цель:** Определить структуру БД (таблицы, столбцы, типы данных, индексы, ограничения) в виде кода **до** создания физической схемы.


**Подходы и инструменты:**

1.  **DDL-скрипты в Git:** Простой подход — хранить `.sql` файлы с `CREATE TABLE` в репозитории. Каждое изменение — это новый миграционный скрипт.

2.  **ORM-спецификации:** Некоторые ORM позволяют описывать схему в коде (например, в Python с SQLAlchemy, в Go с GORM).

3.  **Спецификации в YAML/JSON:** Использовать формат, похожий на OpenAPI, для описания схемы. Например:
```yaml
version: 1.0.0
tables:
  users:
    columns:
      - name: id
        type: BIGINT
        primaryKey: true
        autoIncrement: true
      - name: email
        type: VARCHAR(255)
        unique: true
        notNull: true
      - name: created_at
        type: TIMESTAMPTZ
        defaultValue: now()
    indexes:
      - name: idx_users_created
        columns: [created_at]
        type: btree
```

4.  **Инструменты миграций:**
    *   **[goose](https://github.com/pressly/goose)** (Go): Управление миграциями БД.
    *   **[flyway](https://flywaydb.org/)** (Java, поддерживает SQL и Java): Надёжное применение DDL-скриптов.
    *   **[Liquibase](https://www.liquibase.org/)**: Поддерживает XML, YAML, JSON, SQL.


**Преимущества:**
*   **Версионирование:** Вся история схемы хранится в Git.
*   **Согласованность:** Все среды (dev, test, prod) имеют одинаковую схему.
*   **Тестирование:** Легко развернуть чистую схему для интеграционных тестов.
*   **Аудит:** Прозрачные изменения с комментариями.


---


## 7. Заключение

Spec-First — это универсальный подход, применимый ко всем аспектам системы:

*   **API:** OpenAPI (REST), AsyncAPI (Event-driven).
*   **CLI:** Структурированные спецификации + генерация.
*   **Message Brokers:** AsyncAPI.
*   **Базы данных:** Schema-as-Code и инструменты миграций.


Он превращает проектирование в централизованный, автоматизированный процесс. Несмотря на начальные затраты, Spec-First обеспечивает **предсказуемость, надёжность и масштабируемость** в долгосрочной перспективе, делая его незаменимым для профессиональной разработки.