package config

type Manager struct {
	repoTree confFolder
}

type confElem interface {
	repoList() []*confRepo
	GetName() string
	getParent() *confFolder
	getPath() string
	getPathArray() []string
	GetAllSettings() map[string]interface{}
	getSetting(key string) (interface{}, bool)
}
