package leap

const testVersion = 2

// Provides Boolean leap year test
func IsLeapYear(year int) bool {
	leapYear := false

	if year%4 == 0 && year%100 != 0 {
		leapYear = true
	}

	if year%400 == 0 {
		leapYear = true
	}

	return leapYear
}
