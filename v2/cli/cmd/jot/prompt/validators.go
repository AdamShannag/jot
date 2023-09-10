package prompt

import (
	"errors"
	"regexp"
)

func invalidStringValidator(input string) error {
	if input == "" {
		return errors.New("[Input is empty]")
	}
	if len(input) > 30 {
		return errors.New("[Input is too long]")
	}
	matched, err := regexp.Match("^([a-z0-9]+-)*[a-z0-9]+$", []byte(input))
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("[Invalid input]")
	}
	return nil
}

func invalidNumberValidator(input string) error {
	if input == "" {
		return errors.New("[Input is empty]")
	}
	if len(input) > 9 {
		return errors.New("[Input is too long]")
	}
	matched, err := regexp.Match("^[0-9]*$", []byte(input))
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("[Invalid input]")
	}
	return nil
}

func invalidProjectPathValidator(input string) error {
	if input == "" {
		return errors.New("[Input is empty]")
	}
	return nil
}
