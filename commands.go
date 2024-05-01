package main

type CliCommand struct {
	Name     string
	Desc     string
	Callback func()
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:     "help",
			Desc:     "Print the help menu",
			Callback: CallbackHelp,
		},
		"handp": {
			Name:     "handp",
			Desc:     "Print every card in your hand",
			Callback: CallbackHandPrint,
		},
		"exit": {
			Name:     "exit",
			Desc:     "Terminate the programm immediately",
			Callback: CallbackExit,
		},
		"select": {
			Name:     "select",
			Desc:     "Select the left most card in your hand",
			Callback: CallbackSelect,
		},
		"next": {
			Name:     "next",
			Desc:     "Select the next card in your hand",
			Callback: CallbackNext,
		},
		"prev": {
			Name:     "prev",
			Desc:     "Select the prev card in your hand",
			Callback: CallbackPrev,
		},
		"play": {
			Name:     "play",
			Desc:     "Play the selected card and add it to the board",
			Callback: CallbackPlayCard,
		},
	}
}
