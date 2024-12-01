// config/read.go
package config

import (
	"encoding/json"
	"io/ioutil"
)

func ReadConfig(filePath string) (Config, error) {
	var cfg Config
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
