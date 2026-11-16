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
		return 0, 0, errors.New("Ошибка: в слайсе содержится не два значения")
	}
	stepsInt, err := strconv.Atoi(dataString[0])
	if err != nil {
		return 0, 0, errors.New("Ошибка: Не удалось конвертировать строку в число")
	}
	if stepsInt <= 0 {
		return 0, 0, errors.New("Ошибка: 0 шагов")
	}
	duration, err := time.ParseDuration(dataString[1])
	if err != nil {
		return 0, 0, errors.New("Ошибка: Не удалось конвертировать строку во время")
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
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
}
