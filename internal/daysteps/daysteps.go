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
	dataString := strings.Split(data, ",")
	if len(dataString) != 2 {
		return 0, 0, errors.New("Fail:an insufficient quantity ")
	}
	stepsInt, err := strconv.Atoi(dataString[0])
	if err != nil {
		return 0, 0, errors.New("Fail: conversion string to int")
	}
	if stepsInt <= 0 {
		return 0, 0, errors.New("Fail: steps 0 ")
	}
	duration, err := time.ParseDuration(dataString[1])
	if err != nil {
		return 0, 0, errors.New("Fail: conversion string to time")
	}
	if duration <= 0 {
		return 0, 0, errors.New("Fail: duration 0 ")
	}
	return stepsInt, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceM := float64(steps) * stepLength
	distanceKm := distanceM / float64(mInKm)

	calories, calorieserr := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if calorieserr != nil {
		fmt.Println(calorieserr)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДиcтaнция составила %.2f км.\nBы сожгли %.2f ккал.\n", steps, distanceKm, calories)
}
