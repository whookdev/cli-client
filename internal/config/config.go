package config

type Config struct {
	ConductorUrl string
}

// TODO: make this load from file
func Load() *Config {
	return &Config{
		ConductorUrl: "http://api.localhost:6969/relay",
	}
}
