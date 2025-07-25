package validation

import (
	"errors"
	"regexp"

	"github.com/chera-mihiretu/IKnow/domain/models"
)

func RegisterValidationEmail(user models.User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.UniversityID == nil || user.UniversityID.Hex() == "" {
		return errors.New("university ID is required")
	}

	if user.SchoolID == nil || user.SchoolID.Hex() == "" {
		return errors.New("school ID is required")
	}

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return errors.New("invalid email format")
	}

	if user.PasswordHash == "" {
		return errors.New("password is required")
	}
	if len(user.PasswordHash) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	if user.AcedemicYear <= 0 && user.AcedemicYear >= 6 {
		return errors.New("acedemic year must be between 1 and 5")
	}

	return nil

}
