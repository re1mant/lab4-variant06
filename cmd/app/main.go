package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/re1mant/lab4-variant06/pkg/waterbill"
)

func main() {
	// Подготовка цветного вывода
	headerStyle := color.New(color.FgCyan, color.Bold)
	successStyle := color.New(color.FgGreen)
	warningStyle := color.New(color.FgYellow)
	errorStyle := color.New(color.FgRed, color.Bold)

	headerStyle.Println("=== СИСТЕМА УЧЕТА ВОДОСНАБЖЕНИЯ ===")

	// 1. Вызов F1: WaterUsage
	prevReadings := 120.50
	currReadings := 145.80

	usage, err := waterbill.WaterUsage(prevReadings, currReadings)
	if err != nil {
		errorStyle.Printf("Ошибка при расчете расхода воды: %v\n", err)
		return
	}
	fmt.Printf("Предыдущие показания: %.2f м³\n", prevReadings)
	fmt.Printf("Текущие показания:    %.2f м³\n", currReadings)
	fmt.Printf("Расход воды:          %.2f м³\n\n", usage)

	// 2. Расчет базовой стоимости
	tariff := 185.50 // ₸ за 1 м³
	cost, err := waterbill.WaterCost(usage, tariff)
	if err != nil {
		errorStyle.Printf("Ошибка при расчете стоимости: %v\n", err)
		return
	}
	fmt.Printf("Тариф за 1 м³: %.2f ₸\n", tariff)
	fmt.Printf("Базовая стоимость: %.2f ₸\n\n", cost)

	// 3. Вызов F2: ApplyPenalty (работа с указателем)
	penaltyPercent := 5.0
	err = waterbill.ApplyPenalty(&cost, penaltyPercent)
	if err != nil {
		errorStyle.Printf("Ошибка при начислении штрафа: %v\n", err)
		return
	}
	warningStyle.Printf("Начислен штраф за просрочку: %.1f%%\n", penaltyPercent)
	fmt.Printf("Итоговая стоимость с учетом штрафа: %.2f ₸\n\n", cost)

	// 4. Вызов F3: FormatWaterReport
	owner := "Пичитаев Н.М."
	report, err := waterbill.FormatWaterReport(owner, usage, cost)
	if err != nil {
		errorStyle.Printf("Ошибка при формировании отчета: %v\n", err)
		return
	}

	// 5. Использование пакета uuid из GitHub
	receiptID := uuid.New()

	successStyle.Println("=== СФОРМИРОВАННЫЙ ДОКУМЕНТ ===")
	fmt.Printf("ID Квитанции: %s\n", receiptID.String())
	fmt.Println(report)

	// Проверка корректной перехвати ошибки
	fmt.Println("\n--- Проверка обработки ошибок ---")
	_, errInvalid := waterbill.WaterUsage(100.0, 50.0)
	if errInvalid != nil {
		errorStyle.Printf("Успешный перехват ошибки: %v\n", errInvalid)
	}
}