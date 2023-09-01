package prompt

import (
	"log"

	"github.com/manifoldco/promptui"
)

func (p *PrompterImpl) Prompt(label string, validator Validator) string {
	prop := promptui.Prompt{
		Label:    label,
		Validate: promptui.ValidateFunc(validator),
	}
	prop.HideEntered = true

	input, err := prop.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	return input
}
