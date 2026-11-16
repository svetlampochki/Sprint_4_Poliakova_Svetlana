package spentcalories

import (
	"errors"
	"fmt"
	"log"
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
	DataString := strings.Split(data, ",")
	if len(DataString) != 3 {
		return 0, "", 0, errors.New("Ошибка: в слайсе содержится не три значения")
	}
	steps, err := strconv.Atoi(DataString[0])
	if err != nil {
		return 0, "", 0, errors.New("Ошибка: Не удалось конвертировать строку в число")
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("Ошибка: 0 шагов")
	}
	duration, err := time.ParseDuration(DataString[2])
	if err != nil {
		return 0, "", 0, errors.New("Ошибка: Не удалось конвертировать строку во время")
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("Ошибка: 0 шагов")
	}
	return steps, DataString[1], duration, nil
}

func distance(steps int, height float64) float64 {
	distanceKm := ((height * stepLengthCoefficient) * float64(steps)) / float64(mInKm)
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	Distance := distance(steps, height)
	return Distance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println("Fail: parsing", err)
		return "", err
	}
	if steps <= 0 {
		err := errors.New("Fail: 0 steps")
		log.Println("Fail:", err)
		return "", err
	}
	if weight <= 0 {
		err := errors.New("Fail: weight = 0 ")
		log.Println("Fail:", err)
		return "", err
	}
	if height <= 0 {
		err := errors.New("Fail: height = 0 ")
		log.Println("Fail:", err)
		return "", err
	}
	if duration <= 0 {
		err := errors.New("Fail: duration = 0 ")
		log.Println("Fail:", err)
		return "", err
	}
	var call float64
	var errCall error

	distanceKm := distance(steps, height)
	meanSpeedKmH := meanSpeed(steps, height, duration)
	switch activity {
	case "Ходьба":
		call, errCall = WalkingSpentCalories(steps, weight, height, duration)
	case "Бег":
		call, errCall = RunningSpentCalories(steps, weight, height, duration)
	default:
		err := fmt.Errorf("неизвестный тип тренировки")
		log.Println("Fail:", err)
		return "", err
	}
	if errCall != nil {
		log.Println("Fail spent calories:", errCall)
		return "", errCall
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activity, duration.Hours(), distanceKm, meanSpeedKmH, call)
	return result, nil
}
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Fail: 0 шагов")
	}
	if weight <= 0 {
		return 0, errors.New("Fail: weight = 0 ")
	}
	if height <= 0 {
		return 0, errors.New("Fail: height = 0 ")
	}
	if duration <= 0 {
		return 0, errors.New("Fail: duration = 0 ")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	SpentCalories := (weight * meanSpeed * durationInMinutes) / minInH
	return SpentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("Fail: 0 steps")
	}
	if weight <= 0 {
		return 0, errors.New("Fail: weight = 0 ")
	}
	if height <= 0 {
		return 0, errors.New("Fail: height = 0 ")
	}
	if duration <= 0 {
		return 0, errors.New("Fail: duration = 0 ")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	SpentCalories := (weight * meanSpeed * durationMin) / minInH
	WalkingSpentCalories := SpentCalories * walkingCaloriesCoefficient
	return WalkingSpentCalories, nil
}
