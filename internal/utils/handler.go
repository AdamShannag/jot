package utils

import (
	"errors"
)

type Handler interface {
	HandlePath(path *string) error
	HandleName(name *string) error
}

type DefaultHandler struct {
	formatter Formatter
	validator Validator
}

func NewHandler(formatter Formatter, validator Validator) DefaultHandler {
	return DefaultHandler{
		formatter: formatter,
		validator: validator,
	}
}

func (h *DefaultHandler) HandleName(name *string) error {
	h.formatter.FormatName(name)
	if !h.validator.ValidateName(*name) {
		return errors.New("invalid name project")
	}
	return nil
}

func (h *DefaultHandler) HandlePath(path *string) error {
	h.formatter.FormatPath(path)
	if !h.validator.ValidatePath(*path) {
		return errors.New("invalid path project")
	}
	return nil
}
