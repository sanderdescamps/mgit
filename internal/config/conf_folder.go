package config

import (
	"fmt"
	"strings"
)

type confFolder struct {
	confElem
	parent        *confFolder
	childs        []*confElem
	LocalSettings map[string]interface{}
	Name          string
}

func newconfFolder(name string, parent *confFolder, settings map[string]interface{}) *confFolder {
	if settings == nil {
		settings = map[string]interface{}{}
	}
	return &confFolder{
		Name:          name,
		parent:        parent,
		LocalSettings: settings,
	}
}

func (r confFolder) getParent() *confFolder {
	return r.parent
}

func (r confFolder) GetName() string {
	if val, ok := r.LocalSettings[SETTINGS_GIT_URL_REPO_NAME]; ok {
		return fmt.Sprintf("%v", val)
	}
	return r.Name
}

func (r confFolder) getPathArray() []string {
	if r.parent != nil {
		return append(r.parent.getPathArray(), r.GetName())
	} else if r.GetName() != "" {
		return []string{r.GetName()}
	} else {
		return []string{}
	}
}

func (r confFolder) getPath() string {
	return strings.Join(r.getPathArray(), "/")
}

func (r confFolder) repoList() []*confRepo {
	var repos []*confRepo
	for _, c := range r.childs {
		rl := (*c).repoList()
		if len(rl) > 0 {
			repos = append(repos, rl...)
		}
	}
	return repos
}

func (r confFolder) GetSetting(key string) (interface{}, bool) {
	if val, ok := r.LocalSettings[key]; ok {
		return val, true
	} else if r.getParent() != nil {
		if val, ok := r.getParent().GetSetting(key); ok {
			return val, true
		}
	}
	return nil, false
}

func (r confFolder) GetAllSettings() map[string]interface{} {
	result := r.LocalSettings
	if r.parent != nil {
		for k, v := range r.parent.GetAllSettings() {
			if _, in := result[k]; !in {
				result[k] = v
			}
		}
	}
	return result
}

func (r *confFolder) AddSetting(key string, val interface{}) {
	r.LocalSettings[key] = val
}

func (r *confFolder) AddSettings(settings map[string]interface{}) {
	for k, v := range settings {
		r.AddSetting(k, v)
	}
}

func (r *confFolder) AddChild(child *confElem) {
	r.childs = append(r.childs, child)
}
