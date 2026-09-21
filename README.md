Лабораторная работа №4 — Вариант 6 (waterbill)

Репозиторий содержит выполнение лабораторной работы №4 по дисциплине разработки на Go. Проект представляет собой собственный пакет для расчета расходов на водоснабжение, применения штрафных санкций за просрочку платежа и формирования отчетов.

📌 Задание (Вариант 6: waterbill)

В пакете waterbill реализованы следующие функции:

WaterUsage(prev, curr float64) (float64, error) — вычисление объема израсходованной воды в м³.

WaterCost(cubic, tariff float64) (float64, error) — вычисление базовой стоимости водоснабжения.

ApplyPenalty(cost *float64, penaltyPercent float64) error — изменение стоимости с учетом штрафного процента (работа с указателем).

FormatWaterReport(owner string, cubic, cost float64) (string, error) — формирование форматированной строки отчета.

📁 Структура проекта

lab4-variant06/
├── go.mod
├── go.sum
├── README.md
├── cmd/
│   └── app/
│       └── main.go         # Точка входа в приложение
└── pkg/
    └── waterbill/
        └── waterbill.go    # Пакет с логикой и документацией


🛠 Зависимости (GitHub)

В проекте используются следующие внешние пакеты:

github.com/fatih/color — для форматированного цветного вывода в консоль.

github.com/google/uuid — для генерации уникального идентификатора квитанции.

Установка зависимостей:

go get github.com/fatih/color
go get github.com/google/uuid


🚀 Запуск приложения

Для запуска программы выполните следующую команду из корневой директории проекта:

go run cmd/app/main.go


📖 Проверка документации

Через консоль (go doc):

go doc ./pkg/waterbill
go doc ./pkg/waterbill.WaterUsage


Через веб-интерфейс (godoc):

godoc -http=:6060


После запуска откройте браузер и перейдите по адресу:

http://localhost:6060/pkg/
