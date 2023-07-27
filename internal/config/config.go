package config

type (
	Config struct {
		Redis      Redis      `yaml:"redis"`
		Middleware Middleware `yaml:"middleware"`
	}

	Redis struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Middleware struct {
		UserKey   string `yaml:"userkey"`
		PerMinute int    `yaml:"perminute"`
		Monthly   int    `yaml:"monthly"`
	}
)
