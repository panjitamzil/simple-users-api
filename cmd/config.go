package main

type (
	// config struct
	config struct {
		DB Database `envconfig:"DATABASE"`
	}

	// Database struct
	Database struct {
		Host     string `envconfig:"HOST"`
		Dialect  string `envconfig:"DIALECT"`
		Port     int64  `envconfig:"PORT"`
		Name     string `envconfig:"NAME"`
		Username string `envconfig:"USERNAME"`
		Password string `envconfig:"PASSWORD"`
	}
)
