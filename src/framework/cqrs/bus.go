package cqrs

import (
	"context"
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

func (c *Bus) Handle(ctx context.Context, cmd Command) (any, error) {
	var res any

	for _, v := range c.handlers {
		if v.Name() == cmd.Name() {
			res = v.Handle(ctx, cmd)
			break
		}
	}

	if res == nil {
		return nil, errors.New("handler for" + cmd.Name() + " not found.")
	}

	return res, nil
}
