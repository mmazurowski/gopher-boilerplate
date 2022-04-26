package commands

import log "github.com/sirupsen/logrus"

type SayHelloCommandHandler struct {
	logger *log.Logger
}

func NewSayHelloCommandHandler(logger *log.Logger) SayHelloCommandHandler {
	return SayHelloCommandHandler{logger: logger}
}

func (h SayHelloCommandHandler) Name() string {
	return "SAY_HELLO_COMMAND"
}

func (h SayHelloCommandHandler) Handle(plainCmd any) any {
	cmd := plainCmd.(SayHelloCommand)

	h.logger.Info("Received command " + cmd.Name())

	return "Hello to " + cmd.name
}
