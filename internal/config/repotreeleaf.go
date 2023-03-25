package config

import "strings"

func newRepoTreeLeaf(name string, parent *repoTreeParent, settings map[string]interface{}) repoTreeLeaf {
	return repoTreeLeaf{
		Name:          name,
		parent:        parent,
		LocalSettings: settings,
	}
}

func (r repoTreeLeaf) repoList() []RepoConfig {
	var repos []RepoConfig
	var allSettings map[string]interface{}

	if r.getParent() != nil {
		allSettings = MergeDir(r.getParent().LocalSettings, r.LocalSettings)
	} else {
		allSettings = r.LocalSettings
	}

	newRepo := newRepoConfigFromSettingsMap(allSettings)
	repos = append(repos, *newRepo)
	return repos
}

func (r repoTreeLeaf) getParent() *repoTreeParent {
	return r.parent
}

func (r repoTreeLeaf) getParentPathArray() []string {
	if r.parent != nil {
		return append(r.parent.getParentPathArray(), r.Name)
	} else if r.Name != "" {
		return []string{r.Name}
	} else {
		return []string{}
	}
}

func (r repoTreeLeaf) getParentPath() string {
	return strings.Join(r.getParentPathArray(), "/")
}
