package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("trainings.Parse: len data slice expected: 3, got: %d", len(data))
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("trainings.Parse: invalid steps: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("trainings.Parse: steps must be positive, got %d", steps)
	}

	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("trainings.Parse: invalid duration: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("trainings.Parse: duration must be positive, got %v", duration)
	}

	t.Steps = steps
	t.Duration = duration
	t.TrainingType = data[1]
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distanceInKm := spentenergy.Distance(t.Steps, t.Height)
	avgSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var spentCalories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("trainings.ActionInfo: got error from spentenergy.RunningSpentCalories, error: %v", err)
		}
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("trainings.ActionInfo: got error from spentenergy.WalkingSpentCalories, error: %v", err)
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distanceInKm, avgSpeed, spentCalories,
	), nil
}
