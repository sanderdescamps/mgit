package config

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

type confRepo struct {
	confElem
	parent        *confFolder
	LocalSettings map[string]interface{}
	Name          string
}

func newconfRepo(name string, parent *confFolder, settings map[string]interface{}) confRepo {
	return confRepo{
		Name:          name,
		parent:        parent,
		LocalSettings: settings,
	}
}

func (r confRepo) repoList() []*confRepo {
	return []*confRepo{&r}
}

func (r confRepo) getParent() *confFolder {
	return r.parent
}

func (r confRepo) GetName() string {
	if val, ok := r.GetSettingString(SETTINGS_GIT_URL_REPO_NAME); ok {
		return val
	}
	return r.Name
}

func (r confRepo) getPathArray() []string {
	if r.parent != nil {
		return append(r.parent.getPathArray(), r.GetName())
	} else if r.GetName() != "" {
		return []string{r.GetName()}
	} else {
		return []string{}
	}
}

// Full path of the repo. Including the repo itself.
func (r confRepo) getPath() string {
	sep, _ := r.GetSettingString(SETTINGS_GIT_URL_SEPARATOR)
	return strings.Join(r.getPathArray(), sep)
}

// Recursive parent path
func (r confRepo) getRPath() string {
	var parentPath string
	sep, _ := r.GetSettingString(SETTINGS_GIT_URL_SEPARATOR)
	if val, ok := r.GetSettingString(SETTINGS_GIT_URL_RPATH); ok {
		parentPath = strings.Trim(val, sep)
	} else if r.parent != nil {
		parentPathArray := r.parent.getPathArray()
		parentPath = strings.Join(parentPathArray, sep)
	}
	return parentPath
}

// Fixed parent path
func (r confRepo) getFPath() string {
	var parentPath string
	sep, _ := r.GetSettingString(SETTINGS_GIT_URL_SEPARATOR)
	if val, ok := r.GetSettingString(SETTINGS_GIT_URL_FPATH); ok {
		parentPath = strings.Trim(val, sep)
	}
	return parentPath
}

func (r confRepo) GetSetting(key string) (interface{}, bool) {
	if val, ok := r.LocalSettings[key]; ok {
		return val, true
	} else if r.getParent() != nil {
		if val, ok := r.getParent().getSetting(key); ok {
			return val, true
		}
	}
	if val, ok := default_setting[key]; ok {
		return val, true
	}
	return nil, false
}

func (r confRepo) GetAllSettings() map[string]interface{} {
	result := r.LocalSettings
	if r.getParent() != nil {
		for k, v := range r.getParent().GetAllSettings() {
			if _, in := result[k]; !in {
				result[k] = v
			}
		}
	}
	return result
}

func (r confRepo) GetSettingInt(key string) (int64, bool) {
	val, ok := r.GetSetting(key)
	if !ok {
		return 0, false
	} else if i, isInt := val.(int64); isInt {
		return i, true
	} else if i, err := strconv.ParseInt(fmt.Sprintf("%v", val), 10, 64); err == nil {
		return i, true
	} else {
		return 0, false
	}
}

func (r confRepo) GetSettingString(key string) (string, bool) {
	val, ok := r.GetSetting(key)
	if !ok {
		return "", false
	} else if s, isSting := val.(string); isSting {
		return s, true
	} else {
		return fmt.Sprintf("%v", val), true
	}
}

func (r confRepo) GetSettingBool(key string) (bool, bool) {
	val, ok := r.GetSetting(key)
	if !ok {
		return false, false
	} else if b, isSting := val.(bool); isSting {
		return b, true
	} else if b, err := strconv.ParseBool(fmt.Sprintf("%v", val)); err == nil {
		return b, true
	} else {
		return false, false
	}
}

func (r confRepo) GetSettingFloat(key string) (float64, bool) {
	val, ok := r.GetSetting(key)
	if !ok {
		return 0, false
	} else if f, isFloat := val.(float64); isFloat {
		return f, true
	} else if f, err := strconv.ParseFloat(fmt.Sprintf("%v", val), 64); err == nil {
		return f, true
	} else {
		return 0, false
	}
}

func (r confRepo) GetRepoCloneUrl() string {
	if val, ok := r.GetSettingString(SETTINGS_GIT_URL); ok {
		return val
	} else if val, ok := r.GetSettingString(SETTINGS_GIT_URL_TEMPLATE); ok {
		tplt := template.New("git_url_path")
		tplt, err := tplt.Parse(val)
		if err != nil {
			display.Errorf("failed to parse: %v", val)
			display.Errorf("err: %v", err.Error())
			return ""
		}
		var git_url bytes.Buffer
		tplt.Execute(&git_url, map[string]interface{}{
			"repo_name": r.GetName(),
			"fpath":     r.getFPath(),
			"rpath":     r.getRPath(),
		})
		return git_url.String()
	}
	return ""
}

func (r confRepo) GetFSPath() string {
	root, _ := r.GetSettingString(SETTINGS_GIT_PATH_ROOT)
	root = strings.TrimRight(root, "/")
	path := fmt.Sprintf("%s/%s", root, strings.Join(r.getPathArray(), "/"))
	return pathParse(path)
}
