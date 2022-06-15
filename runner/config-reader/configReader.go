package configreader

import "github.com/deyring/argos/models"

func ReadConfig(filename string) (*models.Config, error) {
	config := &models.Config{}
	err := config.Load(filename)
	if err != nil {
		return nil, err
	}
	return config, nil
}
