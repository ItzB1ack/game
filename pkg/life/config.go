package life

type Config struct {
	Width  int
	Height int
}

func New() *Config {
	cfg := &Config{}
	cfg.Height = 50
	cfg.Width = 100
	return cfg
}
