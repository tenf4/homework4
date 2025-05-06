package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	separatedData := strings.Split(data, ",")
	if len(separatedData) != 3 {
		return 0, "", 0, fmt.Errorf("incorrect data format: %s\n", data)
	}
	steps, err := strconv.Atoi(separatedData[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("steps convertion error: %s\n", err)
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("incorrect steps value: %d\n", steps)
	}
	duration, err := time.ParseDuration(separatedData[2])
	if err != nil {
		fmt.Printf("time parsing error: %s\n", err)
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("incorrect duration value: %d\n", duration)
	}
	return steps, separatedData[1], duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	mean := distanceKm / duration.Hours()
	return mean
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	var distanceKm float64
	var mean float64
	var spentCalories float64
	switch activity {
	case "Ходьба":
		distanceKm = distance(steps, height)
		mean = meanSpeed(steps, height, duration)
		spentCalories, err = WalkingSpentCalories(steps, weight, height, duration)

	case "Бег":
		distanceKm = distance(steps, height)
		mean = meanSpeed(steps, height, duration)
		spentCalories, err = RunningSpentCalories(steps, weight, height, duration)

	default:
		return "", fmt.Errorf("incorrect training type: %s", activity)
	}
	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		activity,
		duration.Hours(),
		distanceKm,
		mean,
		spentCalories,
	), err

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("incorrect steps value\n")
	}
	if weight <= 0 {
		return 0, errors.New("incorrect weight value\n")
	}
	if height <= 0 {
		return 0, errors.New("incorrect height value\n")
	}
	mean := meanSpeed(steps, height, duration)
	durationMins := duration.Minutes()
	spentCalories := (weight * mean * durationMins) / minInH
	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("incorrect steps value\n")
	}
	if weight <= 0 {
		return 0, errors.New("incorrect weight value\n")
	}
	if height <= 0 {
		return 0, errors.New("incorrect height value\n")
	}
	mean := meanSpeed(steps, height, duration)
	durationMins := duration.Minutes()
	spentCalories := (weight * mean * durationMins) / minInH
	return spentCalories * walkingCaloriesCoefficient, nil

}
