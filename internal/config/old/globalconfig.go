package config

// type GlobalConfig struct {
// 	LocalRootPath  string                 `mapstructure:"root_dir"`
// 	GitUrlTemplate string                 `mapstructure:"git_url_template"`
// 	Repos          map[string]interface{} `mapstructure:"repos"`
// }

// func (c *GlobalConfig) ToString() string {
// 	d, err := yaml.Marshal(&c)
// 	if err != nil {
// 		log.Fatalf("error: %v", err)
// 	}
// 	return string(d)
// }

// func (c *GlobalConfig) RepoList() []RepoConfig {
// 	var repos []RepoConfig
// 	flatRepos := flattenRepo(c.Repos, "", nil)
// 	for _, r := range flatRepos {
// 		repoConfig := newRepoConfig()
// 		mapstructure.Decode(r.Settings, &repoConfig)
// 		repoConfig.LocalPath = c.LocalRootPath + "/" + r.LocalRelativePath
// 		if repoConfig.GitUrl == "" {
// 			repoConfig.GitUrl = strings.Replace(c.GitUrlTemplate, "{{ git_path }}", r.LocalRelativePath, -1)
// 		}
// 		repos = append(repos, repoConfig)
// 	}
// 	return repos
// }
