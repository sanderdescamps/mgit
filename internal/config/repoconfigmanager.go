package config

func (cm *RepoConfigManager) LoadRepos(paths ...string) {
	for _, p := range paths {
		c, err := readConfig(p)
		if err != nil {
			display.Error(err.Error())
			continue
		}
		cm.repoTree.loadRepos(c)
	}
}

func (cm *RepoConfigManager) GetRepoList() []RepoConfig {
	return cm.repoTree.repoList()
}
