package model

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
				Method:        "GET",
				Endpoint:      "/pets",
				Responses:     []Response{{Body: listResponse, StatusCode: 200, Label: "success", Headers: nil}},
				Enabled:       true,
			},
			{
				Documentation: "create a pet",
				Method:        "POST",
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
