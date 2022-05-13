package util

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
