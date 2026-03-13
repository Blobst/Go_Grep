package internal

type Commands struct {
	CommandExit    string
	CommandHelp    string
	CommandVersion string
}

func NewCommands() *Commands {
	return &Commands{
		CommandExit:    "{exit",
		CommandHelp:    "{help",
		CommandVersion: "{ver",
	}
}
