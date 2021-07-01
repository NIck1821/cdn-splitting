package configs

type Config struct {
	LogPath  string `toml:"log_path"`
	LogLimit int    `toml:"log_limit"`
}

func NewConfig()  *Config {
	return &Config{
		LogLimit: 50000,
	}
}