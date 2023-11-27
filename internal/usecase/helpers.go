package usecase

import "time"

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
	return len(search) > 0 && len(target) > 0 && len(target) >= len(search) &&
		(containsIgnoreCase(target, search) || containsIgnoreCase(target, search))
}

// containsIgnoreCase checks if the target string contains the search string or not
func containsIgnoreCase(target, search string) bool {
	return len(target) >= len(search) &&
		target[:len(search)] == search ||
		(len(target) > len(search) && containsIgnoreCase(target[1:], search))
}
