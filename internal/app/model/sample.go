package model

const getPetsResponse = `{
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

const createPetResponse = `{
  "id": 1,
  "type": "dog"
}
`

const getPetByIdResponse = `{
  "id": 4,
  "type": "tapir"
}
`

func NewSampleService() Service {

	return Service{
		Version:        "1",
		Name:           "Pet Store Service",
		EndpointPrefix: "/api/v1",
		Port:           3000,
		Hostname:       "0.0.0.0",
		Routes: map[string]Route{
			"getPets": {
				Documentation: "get all pets",
				Method:        "GET",
				Endpoint:      "/pets",
				Responses:     []Response{{Body: getPetsResponse, StatusCode: 200, Label: "success", Headers: nil}},
				Enabled:       true,
			},
			"createPet": {
				Documentation: "create a pet",
				Method:        "POST",
				Endpoint:      "/pets",
				Responses:     []Response{{Body: createPetResponse, StatusCode: 200, Label: "success", Headers: nil}},
				Enabled:       true,
			},
			"getPetById": {
				Documentation: "get pet by id",
				Method:        "GET",
				Endpoint:      "/pets/{petId}",
				Responses:     []Response{{Body: getPetByIdResponse, StatusCode: 200, Label: "success", Headers: nil}},
				Enabled:       true,
			},
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
