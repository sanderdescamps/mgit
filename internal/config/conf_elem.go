package config

type confElem interface {
	repoList() []*confRepo
	GetName() string
	getParent() *confFolder
	getPath() string
	getPathArray() []string
	GetAllSettings() map[string]interface{}
	GetSetting(key string) (interface{}, bool)
	AddSetting(key string, val interface{})
	AddSettings(settings map[string]interface{})
}
