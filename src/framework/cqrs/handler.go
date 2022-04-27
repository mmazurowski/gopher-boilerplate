package cqrs

import "context"

type Handler interface {
	Name() string
	Handle(ctx context.Context, cmd any) any
}
