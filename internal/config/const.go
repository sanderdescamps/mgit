package config

const (
	// name of the repository and overrides the key in the
	SETTINGS_GIT_URL = "git_url"

	// name of the repository and overrides the key in the
	SETTINGS_GIT_URL_REPO_NAME = "git_url_repo_name"

	// Git clone url, supports GO-template
	SETTINGS_GIT_URL_TEMPLATE = "git_url_template"

	// Fixed git path, Used when all repos have the same common prefix
	SETTINGS_GIT_URL_FPATH = "git_url_fpath"

	// Recursive path, concatination of all the parent map keys
	SETTINGS_GIT_URL_RPATH = "git_url_rpath"

	// Recursive path, concatination of all the parent map keys
	SETTINGS_GIT_URL_SEPARATOR = "git_url_sep"

	SETTINGS_CLONE_REPO = "clone"

	// Root path on the filesystem from where the repos will be cloned
	SETTINGS_GIT_PATH_ROOT = "git_root_path"

	// // Local path in filesystem where the repo
	// setting_git_path = "git_path"
)

var default_setting = map[string]interface{}{
	SETTINGS_GIT_URL_SEPARATOR: "/",
	SETTINGS_CLONE_REPO:        true,
	SETTINGS_GIT_PATH_ROOT:     "~/git",
}
