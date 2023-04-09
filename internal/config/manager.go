package config

import (
	"strings"
)

type Manager struct {
	repoTree confElem
}

func (m *Manager) LoadRepos(paths ...string) {
	for _, p := range paths {
		c, err := readConfig(p)
		if err != nil {
			display.Error(err.Error())
			continue
		}
		if _, ok := c["repos"]; ok {
			// m.repoTree.loadRepos(c["repos"].(map[string]interface{}))
			m.repoTree = loadRepos("", nil, c["repos"].(map[string]interface{})).(confFolder)
		} else {
			m.repoTree = loadRepos("", nil, c).(confFolder)
			// m.repoTree.loadRepos(c)
		}

	}
}

func (m *Manager) GetRepoList() []*confRepo {
	return m.repoTree.repoList()
}

func loadRepos(name string, parent *confFolder, repoMap map[string]interface{}) confElem {
	settings := make(map[string]interface{})
	repoNames := []string{}

	for key, val := range repoMap {
		if strings.HasPrefix(key, "_") {
			settings[key[1:]] = val
		} else {
			repoNames = append(repoNames, key)
		}
	}
	var result confElem
	if len(repoNames) > 0 {
		folder := newconfFolder(name, parent, settings)
		for _, rn := range repoNames {
			var child confElem
			if val, ok := repoMap[rn].(map[string]interface{}); ok {
				child = loadRepos(rn, &folder, val)
			} else {
				child = newconfRepo(rn, &folder, nil)
			}

			folder.AddChild(&child)
		}
		result = folder
	} else {
		result = newconfRepo(name, parent, settings)
	}
	return result
}
