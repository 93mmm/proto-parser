# proto-parser
- protosym - консольная утилита для псевдо-парсинга .proto-файлов
### Пример:
Исходный файл: [[examples/example.proto]]
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
Нативно:
```
# 1. Склонировать репозиторий
git clone https://github.com/93mmm/proto-parser.git
cd proto-parser
# 2. Собрать бинарник
make build
# 3. Запуск
./protosym <path-to-file>
# 4. Тестирование
make test
```
С помощью Docker:
```bash
```
### Особенности реализации
- Игнорируются вложенные `message`, `enum`, `oneof` (как указано в ТЗ)
- Разбор syntax, import и других ключевых слов не зависит от форматирования: поддерживаются пробелы, табы и переносы строк
- Написаны юнит и интеграционные тесты с помощью библиотеки
- Реализован `Makefile` для сборки, запуска и тестирования
### Используемые технологии
- Go 1.24
- Docker
- Makefile
- Стандартная библиотека
