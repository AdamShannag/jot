package prompt

import (
	"log"

	"github.com/manifoldco/promptui"
)

func (p *PrompterImpl) Select(label string, items ...string) (int, string) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	index, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return index, result
}
