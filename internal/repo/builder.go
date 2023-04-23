package repo

import "github.com/sanderdescamps/mgit/internal/console"

type Builder struct {
	settings map[string]interface{}
}

func (b *Builder) Url(url string) *Builder {
	b.settings["url"] = url
	return b
}

func (b *Builder) RepoPath(path string) *Builder {
	b.settings["repo_path"] = path
	return b
}

func (b *Builder) Insecure() *Builder {
	b.settings["insecure"] = true
	return b
}

func (b *Builder) Secure() *Builder {
	b.settings["insecure"] = false
	return b
}

func (b *Builder) Display(display console.Display) *Builder {
	b.settings["display"] = display
	return b
}
