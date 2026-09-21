// Package waterbill предоставляет функции для расчета потребления воды,
// стоимости водоснабжения, начисления штрафов и формирования отчетов.
package waterbill

import (
	"errors"
	"fmt"
)

// WaterUsage рассчитывает объем израсходованной воды в кубических метрах
// на основе предыдущих (prev) и текущих (curr) показаний счетчика.
// Возвращает ошибку, если показания отрицательные или текущие меньше предыдущих.
func WaterUsage(prev, curr float64) (float64, error) {
	if prev < 0 || curr < 0 {
		return 0, errors.New("показания счетчика не могут быть отрицательными")
	}
	if curr < prev {
		return 0, fmt.Errorf("текущие показания (%.2f) не могут быть меньше предыдущих (%.2f)", curr, prev)
	}
	return curr - prev, nil
}

// WaterCost рассчитывает базовую стоимость воды исходя из объема в кубических метрах (cubic)
// и тарифа за один кубический метр (tariff).
// Возвращает ошибку при отрицательных входных значениях.
func WaterCost(cubic, tariff float64) (float64, error) {
	if cubic < 0 {
		return 0, fmt.Errorf("объем воды не может быть отрицательным: %.2f", cubic)
	}
	if tariff < 0 {
		return 0, fmt.Errorf("тариф не может быть отрицательным: %.2f", tariff)
	}
	return cubic * tariff, nil
}

// ApplyPenalty изменяет итоговую стоимость (cost) с учетом штрафного процента (penaltyPercent).
// Работает через указатель на переменную и меняет значение напрямую.
func ApplyPenalty(cost *float64, penaltyPercent float64) error {
	if cost == nil {
		return errors.New("указатель на стоимость не может быть nil")
	}
	if *cost < 0 {
		return fmt.Errorf("начальная стоимость не может быть отрицательной: %.2f", *cost)
	}
	if penaltyPercent < 0 || penaltyPercent > 100 {
		return fmt.Errorf("некорректный процент штрафа: %.2f%% (допустимо от 0 до 100)", penaltyPercent)
	}

	*cost += *cost * (penaltyPercent / 100.0)
	return nil
}

// FormatWaterReport формирует итоговую строку отчета о расходах на воду
// для владельца (owner) с форматированием через fmt.Sprintf.
func FormatWaterReport(owner string, cubic, cost float64) (string, error) {
	if owner == "" {
		return "", errors.New("имя владельца не может быть пустым")
	}
	if cubic < 0 || cost < 0 {
		return "", fmt.Errorf("объем (%.2f) и стоимость (%.2f) не могут быть отрицательными", cubic, cost)
	}

	report := fmt.Sprintf("Отчет по водоснабжению | Владелец: %-15s | Расход: %8.2f м³ | К оплате: %10.2f ₸", owner, cubic, cost)
	return report, nil
}