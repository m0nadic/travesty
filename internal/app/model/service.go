package model

type Service struct {
	Version        string            `yaml:"version"`
	Name           string            `yaml:"name"`
	EndpointPrefix string            `yaml:"prefix"`
	Port           int               `yaml:"port"`
	Hostname       string            `yaml:"hostname"`
	Routes         []Route           `yaml:"routes"`
	Headers        map[string]string `yaml:"headers"`
}

const listResponse = `{
  "data": [
    {
      "id": 1,
      "type": "dog"
    },
    {
      "id": 2,
      "type": "cat"
    }
  ]
}
`

const createResponse = `{
  "id": 1,
  "type": "dog"
}
`

func NewSampleService() Service {

	return Service{
		Version:        "1",
		Name:           "Sample Service",
		EndpointPrefix: "/api/v1",
		Port:           3000,
		Hostname:       "0.0.0.0",
		Routes: []Route{
			{
				Documentation: "get all pets",
				Method:        "get",
				Endpoint:      "/pets",
				Responses:     []Response{{Body: listResponse, StatusCode: 200, Label: "success", Headers: nil}},
				Enabled:       true,
			},
			{
				Documentation: "create a pet",
				Method:        "post",
				Endpoint:      "/pets",
				Responses:     []Response{{Body: createResponse, StatusCode: 200, Label: "success", Headers: nil}},
				Enabled:       true,
			},
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
