package config

func (m *Manager) LoadRepos(paths ...string) {
	for _, p := range paths {
		c, err := readConfig(p)
		if err != nil {
			display.Error(err.Error())
			continue
		}
		m.repoTree.loadRepos(c)
	}
}

func (m *Manager) GetRepoList() []*confRepo {
	return m.repoTree.repoList()
}
