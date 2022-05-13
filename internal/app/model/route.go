package model

type Route struct {
	Documentation string     `yaml:"documentation"`
	Method        string     `yaml:"method"`
	Endpoint      string     `yaml:"endpoint"`
	Responses     []Response `yaml:"responses"`
	Enabled       bool       `yaml:"enabled"`
}
