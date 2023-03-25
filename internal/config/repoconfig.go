package config

import (
	"github.com/mitchellh/mapstructure"
)

func newRepoConfig() *RepoConfig {
	return &RepoConfig{
		Clone: true,
	}
}

func newRepoConfigFromSettingsMap(settings map[string]interface{}) *RepoConfig {
	repoConfig := *newRepoConfig()
	err := mapstructure.Decode(settings, &repoConfig)
	if err != nil {
		display.Errorf("could not parse settings to RepoConfig: (%s)", mapToOneLineString(settings))
	}
	return &repoConfig
}

func (r *RepoConfig) addSystemConfig(config map[string]interface{}) {

}

// get git path to clone the repo
func (r *RepoConfig) getRepoPath(config map[string]interface{}) {
}

// func (r *RepoConfig) PrettyLocalPath() string {
// 	usr, _ := user.Current()
// 	path := r.LocalPath
// 	if path == "~" {
// 		path = usr.HomeDir
// 	} else if strings.HasPrefix(path, "~/") {
// 		path = filepath.Join(usr.HomeDir, path[2:])
// 	}
// 	path = strings.Replace(path, "/./", "/", -1)
// 	return path
// }
