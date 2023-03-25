package config

// func getConfigPaths() ([]string, error) {
// 	paths := []string{}
// 	if runtime.GOOS == "linux" {
// 		paths = append(paths,
// 			"etc/mgit/config.yml",
// 			"etc/mgit/config.yaml",
// 			"etc/mgit/config.json",
// 		)
// 	} else if runtime.GOOS == "windows" {
// 		paths = append(paths,
// 			"%ProgramData%/mgit/config.yml",
// 			"%ProgramData%/mgit/config.yaml",
// 			"%ProgramData%/mgit/config.json",
// 		)
// 	} else {
// 		return nil, fmt.Errorf("system configuration is not yet supported on %s", runtime.GOOS)
// 	}

// 	if homeDir, err := os.UserHomeDir(); err == nil {
// 		paths = append(paths,
// 			path.Join(homeDir, ".mgit", "config.yml"),
// 			path.Join(homeDir, ".mgit", "config.yaml"),
// 			path.Join(homeDir, ".mgit", "config.json"),
// 		)
// 	} else {
// 		return nil, fmt.Errorf("failed to get home directory")
// 	}

// 	if pwd, err := os.Getwd(); err == nil {
// 		paths = append(paths,
// 			path.Join(pwd, "config.yml"),
// 			path.Join(pwd, "config.yaml"),
// 			path.Join(pwd, "config.json"),
// 		)
// 	} else {
// 		return nil, fmt.Errorf("failed to get current work directory")
// 	}

// 	return paths, nil
// }

// func getRepoPaths() ([]string, error) {
// 	pwd, err := os.Getwd()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get current work directory")
// 	}
// 	paths := []string{}

// 	paths = append(paths,
// 		path.Join(pwd, "repo.yml"),
// 		path.Join(pwd, "repo.yaml"),
// 		path.Join(pwd, "repo.json"),
// 	)
// 	return paths, nil
// }
