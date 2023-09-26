package handler

import (
	"crypto/rand"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func validateRegistrationInput(ph string, name string, pass string) []string {

	errors := []string{}
	// phone number validation
	if !strings.HasPrefix(ph, "+62") {
		errors = append(errors, "Phone number must start with +62")
	}
	if len(ph) < 10 || len(ph) > 13 {
		errors = append(errors, "Phone number must be at minimum 10 characters and maximum 13 characters")
	}

	// full name validation
	if len(name) < 3 || len(name) > 60 {
		errors = append(errors, "Full name must be at minimum 3 characters and maximum 60 characters")
	}

	// password validation
	if len(pass) < 6 || len(pass) > 64 {
		errors = append(errors, "Password must be at minimum 6 characters and maximum 64 characters")
	}

	hasCapitalChar, hasNumber, hasSpecialChar := false, false, false
	for _, r := range pass {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				hasCapitalChar = true
			}
		} else if unicode.IsNumber(r) {
			hasNumber = true
		} else {
			hasSpecialChar = true
		}
	}
	if !hasCapitalChar || !hasNumber || !hasSpecialChar {
		errors = append(errors, "Password must contain at least 1 capital character AND 1 number AND 1 special (non-alphanumeric) character")
	}

	return errors
}

func generateSalt() ([]byte, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	salt := fmt.Sprintf("%x", b)
	return []byte(salt), nil
}

func hashAndSaltPassword(password string) ([]byte, []byte, error) {
	salt, err := generateSalt()
	if err != nil {
		return nil, nil, err
	}

	// Hash the password with the generated salt
	toHash := append([]byte(password), salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(toHash, bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}

	return hashedPassword, salt, nil

}

func validateUpdateProfileInput(ph string, name string) error {

	if ph == "" && name == "" {
		return fmt.Errorf("phone number or full name must be present in request body param")
	}

	// phone number validation
	if ph != "" {
		if !strings.HasPrefix(ph, "+62") {
			return fmt.Errorf("Phone number must start with +62")
		}
		if len(ph) < 10 || len(ph) > 13 {
			return fmt.Errorf("Phone number must be at minimum 10 characters and maximum 13 characters")
		}

	}

	// full name validation
	if name != "" {
		if len(name) < 3 || len(name) > 60 {
			return fmt.Errorf("Full name must be at minimum 3 characters and maximum 60 characters")
		}
	}

	return nil
}
