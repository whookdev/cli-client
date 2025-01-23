package projects

import "errors"

var (
	ErrEmptyProjectName = errors.New("project name cannot be empty")
)

type InputValidator struct {
	MaxNameLength int
}

func NewInputValidator() *InputValidator {
	return &InputValidator{
		MaxNameLength: 50, // default max length
	}
}

func (v *InputValidator) ValidateName(name string) error {
	if name == "" {
		return ErrEmptyProjectName
	}
	if len(name) > v.MaxNameLength {
		return errors.New("project name too long")
	}
	return nil
}
