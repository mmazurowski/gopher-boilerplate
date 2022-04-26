package cqrs

type Handler interface {
	Name() string
	Handle(cmd any) any
}
