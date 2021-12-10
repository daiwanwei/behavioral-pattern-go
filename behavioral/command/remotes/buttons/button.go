package buttons

import "pattern-golang/behavioral/command/remotes/commands"

type Button struct {
	Command commands.Command
}

func (button *Button) Press() {
	button.Command.Execute()
}
