package cqrs

type Command interface {
	Name() string
}
