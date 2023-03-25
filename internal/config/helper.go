package config

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"
)

func mergeDir(maps ...map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, m := range maps {
		for key, val := range m {
			switch v := val.(type) {
			case bool:
				merged[key] = v
			case int:
				merged[key] = v
			case string:
				merged[key] = v
			case float32, float64:
				merged[key] = v.(float64)
			case map[string]interface{}:
				if merged[key] != nil {
					merged[key] = mergeDir(merged[key].(map[string]interface{}), val.(map[string]interface{}))
				} else {
					merged[key] = val.(map[string]interface{})
				}

			case []string:
				if _, ok := merged[key]; ok {
					merged[key] = append(merged[key].([]string), val.([]string)...)
				} else {
					merged[key] = val.([]string)
				}
			default:
				fmt.Printf("Unsupported value; Can not merge maps; key=%s value=%v(%T)\n", key, val, val)
			}
		}
	}
	return merged
}

// func flattenRepo(repoMap map[string]interface{}, parent string, parentConfig map[string]interface{}) []FlatRepoConfig {
// 	var flatRepoList []FlatRepoConfig
// 	repoSettings := make(map[string]interface{})
// 	var repoNames []string

// 	for key, val := range repoMap {
// 		if strings.HasPrefix(key, "_") {
// 			repoSettings[key[1:]] = val
// 		} else {
// 			repoNames = append(repoNames, key)
// 		}
// 	}

// 	if len(repoNames) == 0 {
// 		flatRepoList = append(flatRepoList, FlatRepoConfig{
// 			RepoName:          parent,
// 			LocalRelativePath: parent,
// 			Settings:          helper.MergeDir(parentConfig, repoSettings),
// 		})
// 	} else {
// 		for _, rn := range repoNames {
// 			var rpath string
// 			if parent != "" {
// 				rpath = parent + "/" + rn
// 			} else {
// 				rpath = rn
// 			}
// 			switch subRepo := repoMap[rn].(type) {
// 			case map[string]interface{}:
// 				flatRepoList = append(flatRepoList, flattenRepo(subRepo, rpath, repoSettings)...)
// 			default:
// 				flatRepoList = append(flatRepoList, FlatRepoConfig{
// 					RepoName:          rn,
// 					LocalRelativePath: rpath,
// 					Settings:          repoSettings,
// 				})
// 			}
// 		}
// 	}
// 	return flatRepoList
// }

// func mapToOneLineString(m map[string]interface{}) string {
// 	var outputs []string
// 	for k, v := range m {
// 		outputs = append(outputs, fmt.Sprintf("%s=%s", k, v))
// 	}
// 	return strings.Join(outputs, ", ")
// }

func pathParse(path string) string {
	usr, _ := user.Current()
	if path == "~" {
		path = usr.HomeDir
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(usr.HomeDir, path[2:])
	}
	path = strings.Replace(path, "/./", "/", -1)
	return path
}
