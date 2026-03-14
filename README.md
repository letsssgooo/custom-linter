# custom-linter

`custom-linter` - статический анализатор для проверки текстов лог-сообщений в Go-проектах.

Проверяются лог-сообщения, передаваемые в функции логирования из следующих пакетов:
- `log` (стандартная библиотека),
- `log/slog`,
- `go.uber.org/zap` (`Logger` и `SugaredLogger`).

## Что проверяет линтер

Правила настраиваются через YAML-конфиг:
- `lowercase` - сообщение должно начинаться с маленькой буквы;
- `english` - сообщение должно содержать только английские буквы;
- `symbols` - сообщение не должно содержать запрещенные/подозрительные символы;
- `sensitive` - сообщение не должно содержать потенциально чувствительные слова.

## Установка

Установка выполняется через `go install`:

```bash
go install github.com/letsssgooo/custom-linter/cmd/customlinter@latest
```

Убедитесь, что каталог с Go-бинарниками (обычно `$GOBIN` или `$GOPATH/bin`) добавлен в `PATH`.

## Конфигурация

Для запуска линтера можно задать переменную окружения `LINTER_CONFIG` с путем к YAML-файлу конфигурации (в противном случае все флаги конфигурации устанавливаются в true).

Пример `rules.yaml`:

```yaml
lowercase: true
english: true
symbols: true
sensitive: true
```

Пример экспорта переменной:

```bash
export LINTER_CONFIG="/path/to/rules.yaml"
```

## Запуск

Линтер запускается как отдельный инструмент (без `go vet`):

```bash
customlinter ./...
```

Пример запуска для конкретного пакета:

```bash
customlinter ./internal/analyzer
```

## Автоисправление (`-fix`)

Линтер можно запускать с флагом `-fix` для автоматического исправления нарушений:

```bash
customlinter -fix ./...
```

Автоисправление поддерживается для правил:
- `symbols`
- `lowercase`

Для остальных правил линтер показывает диагностики без автоисправления.

## Интеграция с golangci-lint

`custom-linter` совместим с `golangci-lint` и может быть добавлен в конфигурацию как внешний линтер.

Пример `.golangci.yml`:

```yaml
version: "2"
linters:
  default: none
  enable:
    - customlinter
```

## Тесты

Запуск тестов:

```bash
go test ./...
```

Тестовые сценарии анализатора находятся в `internal/analyzer/testdata/src`.

## Структура проекта

- `cmd/customlinter/main.go` - точка входа CLI;
- `internal/analyzer/analyzer.go` - логика анализа и диагностики;
- `internal/config/config.go` - загрузка и схема YAML-конфига;
- `internal/rules/` - движок правил и их реализации;
- `rules.yaml` - пример конфигурации.
