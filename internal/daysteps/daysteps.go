package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	separatedData := strings.Split(data, ",")
	if len(separatedData) != 2 {
		return 0, 0, fmt.Errorf("Неправильный формат данных: %s", data)
	}

	steps, err := strconv.Atoi(separatedData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка конвертации шагов: %s", err)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("Некорректное количество шагов: %d", steps)
	}
	duration, err := time.ParseDuration(separatedData[1])
	if err != nil {
		fmt.Printf("Ошибка в парсинге времени: %s\n", err)
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("Отрицательная или нулевая продолжительность")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("Ошибка парсинга: %s", err)
	}
	if steps <= 0 {
		return ""
	}
	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d\nДистанция составила %f км.\nВы сожгли %f ккал.", steps, distanceKm, spentCalories)
}
