package features

type Features struct {
	value string
}

func (f Features) String() string {
	return f.value
}

func Greeter() Features {
	return Features{value: "greet-feature-enum"}
}
