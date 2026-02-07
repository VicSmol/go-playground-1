# Задача: Реализовать пакет `setx`

Необходимо создать пакет `setx` в Go, который будет предоставлять реализацию множества (Set) на основе map[T]struct{}. В рамках задачи нужно реализовать две основные структуры данных: `SetString` и `SetInt`, а также методы для работы с ними.

---

## 1. **SetString**

Структура данных `SetString` должна быть реализована на основе `map[string]struct{}`.

### Методы SetString:

1. **NewSetString()** - создает новое пустое множество строк.
2. **Add(s string)** - добавляет строку в множество.
3. **Remove(s string)** - удаляет строку из множества.
4. **Contains(s string) bool** - проверяет, содержится ли строка в множестве.
5. **Size() int** - возвращает количество элементов в множестве.
6. **ToSlice() []string** - возвращает все элементы множества в виде среза строк.
7. **IsEmpty() bool** - проверяет, пусто ли множество.

Пример использования:
```go
set := NewSetString()
set.Add("hello")
set.Add("world")
fmt.Println(set.Contains("hello")) // true
fmt.Println(set.Size()) // 2
```

---

## 2. **SetInt**

Структура данных `SetInt` должна быть реализована на основе `map[int]struct{}`.

### Методы SetInt:

1. **NewSetInt()** - создает новое пустое множество целых чисел.
2. **Add(i int)** - добавляет целое число в множество.
3. **Remove(i int)** - удаляет целое число из множества.
4. **Contains(i int) bool** - проверяет, содержится ли целое число в множестве.
5. **Size() int** - возвращает количество элементов в множестве.
6. **ToSlice() []int** - возвращает все элементы множества в виде среза целых чисел.
7. **IsEmpty() bool** - проверяет, пусто ли множество.

Пример использования:
```go
set := NewSetInt()
set.Add(1)
set.Add(2)
fmt.Println(set.Contains(1)) // true
fmt.Println(set.Size()) // 2
```

---

## Требования:

- Пакет должен находиться в директории internal/setx.
- Необходимо реализовать только указанные структуры и методы.
- Не использовать внешние библиотеки, кроме стандартных библиотек Go.
- Использовать `struct{}` в качестве значения в map для экономии памяти.

---

## Пример структуры пакета:

```go
package setx

type SetString struct {
    // поля структуры
}

func NewSetString() *SetString
func (s *SetString) Add(str string)
func (s *SetString) Remove(str string)
func (s *SetString) Contains(str string) bool
func (s *SetString) Size() int
func (s *SetString) ToSlice() []string
func (s *SetString) IsEmpty() bool

type SetInt struct {
    // поля структуры
}

func NewSetInt() *SetInt
func (s *SetInt) Add(i int)
func (s *SetInt) Remove(i int)
func (s *SetInt) Contains(i int) bool
func (s *SetInt) Size() int
func (s *SetInt) ToSlice() []int
func (s *SetInt) IsEmpty() bool
```

Эти структуры и методы должны быть готовы к использованию в других частях проекта.