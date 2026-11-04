package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("daysteps.Parse: expected len datastring 2, got: %d", len(data))
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("daysteps.Parse: invalid steps, error: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("daysteps.Parse: steps must be positive, got %d", steps)
	}

	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("daysteps.Parse: invalid duration, error: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("daysteps.Parse: duration must be positive, got %v", duration)
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distanceInKm := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("daysteps.ActionInfo: %v", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distanceInKm, spentCalories,
	), nil
}
