package helper

import (
	"os/user"
	"path/filepath"
	"strings"
)

func PathParse(path string) string {
	usr, _ := user.Current()
	if path == "~" {
		path = usr.HomeDir
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(usr.HomeDir, path[2:])
	}
	path = strings.Replace(path, "/./", "/", -1)
	return path
}
