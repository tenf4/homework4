package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
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
		return 0, 0, fmt.Errorf("incorrect data format: %s", data)
	}

	steps, err := strconv.Atoi(separatedData[0])
	if err != nil {
		return 0, 0, fmt.Errorf("steps convertion fail: %w", err)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("incorrect steps value: %d", steps)
	}
	duration, err := time.ParseDuration(separatedData[1])
	if err != nil {
		fmt.Printf("time parsing fail: %s\n", err)
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, errors.New("negative or zero time value")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Print(errors.New("parsing error"))
		return ""
	}
	if steps < 1 {
		log.Print(errors.New("incorrect steps value"))
		return ""
	}
	distanceKm := float64(stepLength) * float64(steps) / 1000
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, spentCalories)

}
