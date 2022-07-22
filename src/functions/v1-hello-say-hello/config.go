package main

type config struct {
	appName string
}

func createConfig() *config {
	return &config{appName: "APP_NAME"}
}
