# Преобразования сигналов

## Описание

Данная программа создана для проекта по физике. В ней реализована генерация сигналов колебательного контура, случайных сигналов и шумов. С помощью быстрого преобразования Фурье сигнал раскладывается на частоты, амплитуды и фазы. Далее сигнал можно отфильтровать и восстановить с помощью обратного преобразования Фурье.

## Функционал

  * Генерация сигналов
  * Генерация шумов
  * Быстрое преобразование Фурье
  * Обратное преобразование Фурье
  * Умножение многочленов
  * Фильтрация сигнала
  * Визуализация сигналов

## Файлы

  * `main.go` - пример использования
  * `FFTW.go` - реализация БПФ, ОБПФ, умножения многочленов, визуализации сигналов
  * `Generator.go` - реализация генерации сигналов и шумов
  * `Filters.go` - реализация фильтрации сигнала по амплитуде и по количеству частот
  * `ElectricalCircuit.go` - реализация сигналов колебательного контура
  * `main_test.go` - тесты, где видно, что точность фильтрации по количеству частот в лучшем случае достигает 100%

## Пример фильтрации:

До фильтрации, файл `noise.png`:

![Image alt](https://github.com/KirillSozin/SignalConversions/blob/master/noise.png)

После фильтрации, файл `filtered.png`:

![Image alt](https://github.com/KirillSozin/SignalConversions/blob/master/filtered.png)

## Сборка и запуск

  1. Скачать репозиторий
  2. Перейти в папку с проектом
  3. Выполнить команду `go build -o name`
  4. Выполнить команду `./name`

## Что будет добавлено

  * Коррекция фаз и амплитуд при разложении сигнала на частоты
  * Будет переписан БПФ: он станет итеративным и разделенным на потоки
  * Будет добавлен консольный интерфейс
  * Можно будет сгенерировать сигнал для произвольного колебательного контура

## Лицензия

[MIT](https://choosealicense.com/licenses/mit/)