package dto

import (
	"errors"
	"regexp"
)

// Regular expressions for validation
var (
	emailRegex    = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	passwordRegex = regexp.MustCompile(`^(?=.*[A-Z])(?=.*[!@#$&*])(?=.*[0-9])(?=.*[a-z]).{8,}$`)
)

// Entity roles
type EntityRole string

const (
	Admin EntityRole = "Admin"
	User  EntityRole = "User"
	Guest EntityRole = "Guest"
)

// ValidRoles contains the list of valid roles
var ValidRoles = []EntityRole{Admin, User, Guest}

// RegisterUserDto represents the data needed to register a user
type RegisterUserDto struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      EntityRole
}

// CreateRegisterUserDto validates the input and creates a new RegisterUserDto instance
func CreateRegisterUserDto(props map[string]interface{}) (*RegisterUserDto, error) {
	email, emailOk := props["email"].(string)
	password, passwordOk := props["password"].(string)
	firstName, firstNameOk := props["firstName"].(string)
	lastName, lastNameOk := props["lastName"].(string)
	role, roleOk := props["role"].(EntityRole)

	if !emailOk || !emailRegex.MatchString(email) {
		return nil, errors.New("email property is required")
	}
	if !passwordOk || !passwordRegex.MatchString(password) {
		return nil, errors.New("password must contain at least eight characters, one special character, one uppercase letter, and one number")
	}
	if !firstNameOk {
		return nil, errors.New("firstName property is required")
	}
	if !lastNameOk {
		return nil, errors.New("lastName property is required")
	}
	if !roleOk || !isValidRole(role) {
		return nil, errors.New("invalid role property")
	}

	return &RegisterUserDto{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
	}, nil
}

// isValidRole checks if the role is valid
func isValidRole(role EntityRole) bool {
	for _, r := range ValidRoles {
		if role == r {
			return true
		}
	}
	return false
}
