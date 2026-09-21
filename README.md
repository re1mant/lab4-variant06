<div align="center">

#  Лабораторная работа №4 — Вариант 6 (`waterbill`)

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![GitHub Repo](https://img.shields.io/badge/GitHub-lab4--variant06-181717?style=for-the-badge&logo=github)

**Модуль на Go для расчета расходов на водоснабжение, применения пени и генерации квитанций.**

---

</div>

##  Описание проекта

Данный проект разработан в рамках Лабораторной работы №4. Он содержит собственный пакет `waterbill` для расчета коммунальных платежей за воду, обработку возможных ошибок, примеры работы с указателями и интеграцию со сторонними библиотеками GitHub.

---

##  Структура проекта

```text
lab4-variant06/
├── go.mod                  # Файл конфигурации модуля Go
├── go.sum                  # Контрольные суммы зависимостей
├── README.md               # Документация проекта
├── cmd/
│   └── app/
│       └── main.go         # Точка входа в приложение
└── pkg/
    └── waterbill/
        └── waterbill.go    # Пакет с бизнес-логикой (Вариант 6)
