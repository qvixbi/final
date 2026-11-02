package nextdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// afterNow проверяет, что дата date позже даты now.
// Время суток при сравнении не учитывается.
func afterNow(date, now time.Time) bool {
	d := date.Format("20060102")
	n := now.Format("20060102")

	return d > n
}

// NextDate вычисляет следующую дату задачи.
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Преобразуем исходную дату в time.Time.
	date, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", fmt.Errorf("неверная исходная дата: %w", err)
	}

	// Разбираем правило повторения.
	parts := strings.Fields(repeat)

	if len(parts) == 0 {
		return "", fmt.Errorf("правило повторения не указано")
	}

	switch parts[0] {

	case "d":
		// Проверяем формат правила d.
		if len(parts) != 2 {
			return "", fmt.Errorf("неверный формат правила d")
		}

		interval, err := strconv.Atoi(parts[1])
		if err != nil || interval < 1 || interval > 400 {
			return "", fmt.Errorf("интервал должен быть от 1 до 400 дней")
		}

		// Переносим дату вперёд, пока она не станет позже now.
		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}

	case "y":
		// Для ежегодной задачи уточнения не нужны.
		if len(parts) != 1 {
			return "", fmt.Errorf("неверный формат правила y")
		}

		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}

	default:
		return "", fmt.Errorf("неподдерживаемое правило повторения: %s", parts[0])
	}

	// Возвращаем дату в формате YYYYMMDD.
	return date.Format("20060102"), nil
}
