package internal

type Commands struct {
	CommandExit        string
	CommandHelp        string
	CommandVersion     string
	CommandUpdateCheck string
}

func NewCommands() *Commands {
	return &Commands{
		CommandExit:        "{exit",
		CommandHelp:        "{help",
		CommandVersion:     "{ver",
		CommandUpdateCheck: "{upd",
	}
}
