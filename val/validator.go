package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 25); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("username can only contain letters, numbers and underscores")
	}
	return nil
}

func ValidateFullname(value string) error {
	if err := ValidateString(value, 3, 25); err != nil {
		return err
	}
	if !isValidFullname(value) {
		return fmt.Errorf("fullname can only contain letters and spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 50)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}
