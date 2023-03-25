package config

// func ReadConfigFromPaths(paths ...string) (*GlobalConfig, error) {
// 	var mergedConfig Config
// 	for _, path := range paths {
// 		var config map[string]interface{}
// 		var err error
// 		if strings.HasSuffix(path, ".json") {
// 			config, err = readJson(path)
// 			if err != nil {
// 				return nil, err
// 			}

// 		} else if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
// 			config, err = readYaml(path)
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			return nil, fmt.Errorf("invalid config file: %s", path)
// 		}
// 		mergedConfig = helper.MergeDir(mergedConfig, config)
// 	}
// 	var globalConfig GlobalConfig
// 	err := mapstructure.Decode(mergedConfig, &globalConfig)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &globalConfig, nil
// }
