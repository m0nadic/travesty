package model

type Response struct {
	Body       string        `yaml:"body"`
	StatusCode int           `yaml:"statusCode"`
	Label      string        `yaml:"label"`
	Headers    []interface{} `yaml:"headers"`
}
