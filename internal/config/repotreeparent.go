package config

import (
	"reflect"
	"strings"
)

func newRepoTreeParent(name string, parent *repoTreeParent, settings map[string]interface{}) repoTreeParent {
	if settings == nil {
		settings = map[string]interface{}{}
	}
	return repoTreeParent{
		Name:          name,
		parent:        parent,
		LocalSettings: settings,
	}
}

func (r repoTreeParent) getParent() *repoTreeParent {
	return r.parent
}

func (r repoTreeParent) getParentPathArray() []string {
	if r.parent != nil {
		return append(r.parent.getParentPathArray(), r.Name)
	} else if r.Name != "" {
		return []string{r.Name}
	} else {
		return []string{}
	}
}

func (r repoTreeParent) getParentPath() string {
	return strings.Join(r.getParentPathArray(), "/")
}

func (r repoTreeParent) repoList() []RepoConfig {
	var repos []RepoConfig
	for _, c := range r.childs {
		rl := c.repoList()
		if len(rl) > 0 {
			repos = append(repos, rl...)
		}
	}
	return repos
}

func (r *repoTreeParent) loadRepos(repoMap map[string]interface{}) {
	settings := make(map[string]interface{})
	repoNames := []string{}

	for key, val := range repoMap {
		if strings.HasPrefix(key, "_") {
			settings[key[1:]] = val
		} else {
			repoNames = append(repoNames, key)
		}
	}

	if len(repoNames) > 0 {
		for _, rn := range repoNames {
			isLeaf := true
			repoSettings := make(map[string]interface{})
			if repoMap[rn] != nil && reflect.ValueOf(repoMap[rn]).Kind() == reflect.Map {
				for key, val := range repoMap[rn].(map[string]interface{}) {
					if strings.HasPrefix(key, "_") {
						repoSettings[key[1:]] = val
					} else {
						isLeaf = false
					}
				}
			}
			if isLeaf {
				newLeaf := newRepoTreeLeaf(rn, r, repoSettings)
				r.childs = append(r.childs, newLeaf)
			} else {
				newParent := newRepoTreeParent(rn, r, nil)
				newParent.loadRepos(repoMap[rn].(map[string]interface{}))
				r.childs = append(r.childs, newParent)
			}
		}
	}
}
