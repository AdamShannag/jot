package prompt

import "fmt"

func (p *PrompterImpl) Start(command string) {
	if command != NEW {
		fmt.Println("Invalid command, please use [new]")
		return
	}

	p.new()
}
