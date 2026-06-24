package config

type Target struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Config struct {
	Targets        []Target `yaml:"targets"`
	TimeoutSeconds int      `yaml:"timeout_seconds"`
}

func Load(path string) (*Config, error) {
	return nil, nil
}
