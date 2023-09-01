package prompt

import "errors"

func invalidStringValidator(input string) error {
	if input == "" {
		return errors.New("invalid input")
	}
	return nil
}
