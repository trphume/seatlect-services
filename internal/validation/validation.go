package validation

func ValidUsername(username string) bool {
	if len(username) < 3 && len(username) > 20 {
		return false
	}

	return true
}

func ValidPassword(password string) bool {
	if len(password) < 3 && len(password) > 30 {
		return false
	}

	return true
}

func ValidEmail(email string) bool {
	// TODO: proper regex
	if len(email) < 3 {
		return false
	}

	return true
}
