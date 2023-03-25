package config

// type Config map[string]interface{}

//	type FlatRepoConfig struct {
//		RepoName          string
//		LocalRelativePath string
//		Settings          map[string]interface{}
//	}

type RepoConfig struct {
	LocalPath string
	GitUrl    string `mapstructure:"git_repo"`
	Clone     bool   `mapstructure:"clone"`
}

type RepoConfigManager struct {
	repoTree repoTreeParent
}

type repoTreeNode interface {
	repoList() []RepoConfig
	getParent() *repoTreeParent
	getParentPath() string
	getParentPathArray() []string
}

type repoTreeParent struct {
	repoTreeNode
	parent        *repoTreeParent
	childs        []repoTreeNode
	LocalSettings map[string]interface{}
	Name          string
}

type repoTreeLeaf struct {
	repoTreeNode
	parent        *repoTreeParent
	LocalSettings map[string]interface{}
	Name          string
}
