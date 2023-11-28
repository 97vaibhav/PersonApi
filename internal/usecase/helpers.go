package usecase

import (
	"strings"
	"time"
)

func CalculateAge(birthday string) int {
	birthDate, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0
	}

	age := time.Now().Year() - birthDate.Year()

	// Adjust age if the birthday hasn't occurred yet this year
	if time.Now().Before(time.Date(time.Now().Year(), birthDate.Month(), birthDate.Day(), 0, 0, 0, 0, time.UTC)) {
		age--
	}

	return age
}

func ContainsName(target, search string) bool {
	target = strings.ToLower(target)
	search = strings.ToLower(search)
	return strings.Contains(target, search)
}
