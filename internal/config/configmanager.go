package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

type ConfigManager struct {
	configPaths []string
}

// Add possible config path. Config files need to be added by increasing priority.
func (cm *ConfigManager) AddConfigPath(path string) {
	cm.configPaths = append(cm.configPaths, path)
}

// Add possible config path with multiple extentions. Config files need to be added by increasing priority.
func (cm *ConfigManager) AddConfigPathE(dirPath string, fileName string, ext []string) {
	for _, e := range ext {
		cm.AddConfigPath(path.Join(dirPath, fmt.Sprintf("%s.%s", fileName, e)))
	}
}

// Add system config paths and home dir config path
func (cm *ConfigManager) AddDefaultPaths() {
	if runtime.GOOS == "linux" {
		cm.AddConfigPathE("etc/mgit", "config", []string{"yml", "yaml", "json"})
	} else if runtime.GOOS == "windows" {
		cm.AddConfigPathE("%ProgramData%/mgit", "config", []string{"yml", "yaml", "json"})
	} else {
		display.Warningf("system configuration is not yet supported on %s. Skipping system config", runtime.GOOS)
	}

	if homeDir, err := os.UserHomeDir(); err == nil {
		cm.AddConfigPathE(path.Join(homeDir, ".mgit"), "config", []string{"yml", "yaml", "json"})
	} else {
		display.Warningf("couldn't get home directory. Skipping homedir config.")
	}
}

func (cm *ConfigManager) GetConfigMap() map[string]interface{} {
	var mergedConfig map[string]interface{}
	for _, path := range cm.configPaths {
		var config map[string]interface{}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			display.Debugf("config file (%s) not found. Skip and fo to next", path)
			continue
		}

		var err error
		if strings.HasSuffix(path, ".json") {
			config, err = readJson(path)
			if err != nil {
				display.Errorf("failed to read config file: %s", path)
				continue
			}

		} else if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
			config, err = readYaml(path)
			if err != nil {
				display.Errorf("failed to read config file: %s", path)
				continue
			}
		} else {
			display.Errorf("invalid config file extention: %s", path)
			continue
		}
		mergedConfig = MergeDir(mergedConfig, config)
	}
	return mergedConfig
}

func (cm *ConfigManager) GetConfigs() []map[string]interface{} {
	configs := []map[string]interface{}{}
	for _, path := range cm.configPaths {
		var config map[string]interface{}

		if _, err := os.Stat(path); os.IsNotExist(err) {
			display.Debugf("config file (%s) not found. Skip and fo to next", path)
			continue
		}

		var err error
		if strings.HasSuffix(path, ".json") {
			config, err = readJson(path)
			if err != nil {
				display.Errorf("failed to read config file: %s", path)
				continue
			}

		} else if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
			config, err = readYaml(path)
			if err != nil {
				display.Errorf("failed to read config file: %s", path)
				continue
			}
		} else {
			display.Errorf("invalid config file extention: %s", path)
			continue
		}
		configs = append(configs, config)
	}
	return configs
}
