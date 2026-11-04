package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0, fmt.Errorf("spentenergy.WalkingSpentCalories: duration must be positive, got %v", duration)
	}

	if steps <= 0 {
		return 0, fmt.Errorf("spentenergy.WalkingSpentCalories: steps must be positive, got %d", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("spentenergy.WalkingSpentCalories: weight must be positive, got %.2f", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("spentenergy.WalkingSpentCalories: height must be positive, got %.2f", height)
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	spentCalories := (weight * avgSpeed * duration.Minutes()) / minInH
	spentCalories *= walkingCaloriesCoefficient
	return spentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0, fmt.Errorf("spentenergy.RunningSpentCalories: duration must be positive, got %v", duration)
	}

	if steps <= 0 {
		return 0, fmt.Errorf("spentenergy.RunningSpentCalories: steps must be positive, got %d", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("spentenergy.RunningSpentCalories: weight must be positive, got %.2f", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("spentenergy.RunningSpentCalories: height must be positive, got %.2f", height)
	}

	avgSpeed := MeanSpeed(steps, height, duration)

	spentCalories := (weight * avgSpeed * duration.Minutes()) / minInH
	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0
	}

	distanceInKm := Distance(steps, height)
	avgSpeed := distanceInKm / duration.Hours()
	return avgSpeed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceInKm := (stepLength * float64(steps)) / mInKm
	return distanceInKm
}
