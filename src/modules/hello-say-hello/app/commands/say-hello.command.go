package commands

type SayHelloCommand struct {
	name string
}

func NewSayHelloCommand(name string) SayHelloCommand {
	return SayHelloCommand{name: name}
}

func (c SayHelloCommand) Name() string {
	return "SAY_HELLO_COMMAND"
}
