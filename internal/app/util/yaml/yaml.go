package yaml

import (
	"gopkg.in/yaml.v2"
	"os"
	"travesty/internal/app/model"
)

func GenerateFile(fileName string) error {
	service := model.NewSampleService()

	data, err := yaml.Marshal(&service)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadFile(fileName string) (model.Service, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return model.Service{}, err
	}
	var service model.Service
	err = yaml.Unmarshal(data, &service)
	if err != nil {
		return model.Service{}, err
	}

	return service, nil
}
