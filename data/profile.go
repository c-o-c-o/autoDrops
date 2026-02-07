package data

type Profile struct {
	Client  string  `yaml:"Client"`
	Advance int     `yaml:"Advance"`
	Actors  []Actor `yaml:"Actors"`
}

type Actor struct {
	Layer   int    `yaml:"Layer"`
	Pattern string `yaml:"Pattern"`
	Target  string `yaml:"Target"`
}
