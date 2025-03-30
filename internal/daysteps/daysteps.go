package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format: expected 2 values, got %d, "+
			"data: %q", len(parts), data)
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("step count must be greater than zero")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, trackedTime, err := parsePackage(data)
	if err != nil {
		fmt.Printf("Ошибка обработки данных: %v\n", err)
		return ""
	}
	distance := (float64(steps) * StepLength) / 1000
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, trackedTime)

	summary := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.", steps, distance, calories)
	return summary
}
