# proto-parser
- protosym - консольная утилита для псевдо-парсинга .proto-файлов
### Пример:
Исходный [файл](examples/example.proto)
Запуск:
```bash
./protosym examples/example.proto
```
Вывод:
```
google/protobuf/timestamp.proto import 5:8-41
Example service 9:9-16
ExampleRPC method 10:7-17
ExampleEnum enum 13:6-17
ExampleRPCRequest message 19:9-26
ExampleRPCResponse message 26:9-27
```
### Установка и запуск
Установка:
```bash
git clone https://github.com/93mmm/proto-parser.git
cd proto-parser
```
Запуск нативно:
```
# 0. Установить зависимости
go mod tidy
# 1. Собрать бинарник
make build
# 2. Запуск
./protosym <path-to-file>
# 3. Тестирование
make test
```
Запуск с помощью Docker:
```bash
# 1. Запуск
make docker_run file=<relative-path-to-file>
# 2. Тестирование
make docker_test

# Сбор контейнера происходит с помощью
make docker_build
```
### Особенности реализации
- Игнорируются вложенные `message`, `enum`, `oneof` (как указано в ТЗ)
- Разбор syntax, import и других ключевых слов не зависит от форматирования: поддерживаются пробелы, табы и переносы строк
- Написаны юнит и интеграционные тесты с помощью библиотеки
- Реализован `Makefile` для нативной сборки, запуска и тестирования
- Реализован `Dockerfile` для сборки, запуска и тестирования приложения
### Используемые технологии
- Go 1.24
- Стандартная библиотека
- Docker
- Makefile
- `testify`
