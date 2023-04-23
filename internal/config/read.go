package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"path/filepath"

	"gopkg.in/yaml.v3"
)

func readJson(path string) (map[string]interface{}, error) {
	path = filepath.Clean(path)
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read json file (%s): %s", path, err.Error())
	}

	byteValue, _ := io.ReadAll(jsonFile)
	var data map[string]interface{}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to parse json file (%s): %s", path, err.Error())
	}
	return data, nil
}

func readYaml(path string) (map[string]interface{}, error) {
	path = filepath.Clean(path)
	yamlFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to read YAML file (%s): %s", path, err.Error())
	}

	byteValue, _ := io.ReadAll(yamlFile)
	var data map[string]interface{}
	err = yaml.Unmarshal(byteValue, &data)
	if err != nil {
		return nil, fmt.Errorf("unable to parse YAML file (%s): %s", path, err.Error())
	}
	return data, nil
}

func readConfig(path string) (map[string]interface{}, error) {
	var config map[string]interface{}
	var err error
	if strings.HasSuffix(path, ".json") {
		config, err = readJson(path)
		if err != nil {
			return nil, err
		}

	} else if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
		config, err = readYaml(path)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("invalid config file: %s", path)
	}
	return config, nil
}
