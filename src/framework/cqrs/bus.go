package cqrs

import (
	"errors"
)

type Bus struct {
	handlers []Handler
}

func Factory(handlers []Handler) Bus {
	return Bus{
		handlers: handlers,
	}
}

func (c *Bus) Handle(cmd Command) (any, error) {
	var res any

	for _, v := range c.handlers {
		if v.Name() == cmd.Name() {
			res = v.Handle(cmd)
			break
		}
	}

	if res == nil {
		return nil, errors.New("handler for" + cmd.Name() + " not found.")
	}

	return res, nil
}
