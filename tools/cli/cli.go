package cli

import (
	"fmt"
	"strings"

	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
)

type CLI struct{}

func (c *CLI) execDiePC() {
	statemachine.UseStateMachine().PC().SetState(statemachine.PC_DEAD)
}

func (c *CLI) Run() {
	go func() {
		for {
			var command string
			fmt.Scan(&command)

			switch strings.ToLower(command) {
			case "die_pc":
				c.execDiePC()
			}
		}
	}()
}

func NewCLI() *CLI {
	return new(CLI)
}
