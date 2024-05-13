package password

import (
	"errors"
	"unicode"
)

func ValidatePassword(password string) error {

	countNum := 0
	countSym := 0
	countUpper := 0
	for _, v := range password {
		if unicode.IsNumber(v) {
			countNum++
		}
		if unicode.IsSymbol(v) {
			countSym++
		}
		if unicode.IsUpper(v) {
			countUpper++
		}
	}
	if countNum < 2 {
		return errors.New("the password must contain at least 2 numbers")
	}

	if countUpper < 1 {
		return errors.New("the password must contain at least one capital letter")
	}

	if countSym < 1 {
		return errors.New("the password must contain at least one special character")
	}

	return nil
}
