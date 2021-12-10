package buttons

import "behavioral-pattern-go/command/remotes/commands"

type Button struct {
	Command commands.Command
}

func (button *Button) Press() {
	button.Command.Execute()
}
