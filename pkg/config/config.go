package config

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port string
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			Port: ":8080",
		},
	}
} 