package repo

import (
	"os"

	"github.com/go-git/go-git/v5"
)

func (r *Repo) isClonedOnFileSystem() bool {
	if _, err := os.Stat(r.RepoPath); os.IsNotExist(err) {
		return false
	} else {
		_, err := git.PlainOpen(r.RepoPath)
		if err == nil {
			return true
		}
	}
	return false
}

func (r *Repo) Clone() {
	if r.repo != nil {
		r.Display.Skip("Repo already cloned")
	} else if !r.isClonedOnFileSystem() {
		r.Display.Debugf("Clone repo %s...", r.Url)
		repo, err := git.PlainClone(r.RepoPath, false, &git.CloneOptions{
			URL:      r.Url,
			Progress: r.Display.Writer(),
			// Auth: transport.AuthMethod{},
		})
		if err != nil {
			r.Display.Error(err.Error())
			return
		}
		r.repo = repo
	} else {
		repo, err := git.PlainOpen(r.RepoPath)
		if err == nil {
			r.Display.Okf("Repo already cloned: %s", r.Url)
		} else {
			r.Display.Errorf("failed to load repo: %s\n%v", r.Url, err.Error())
			return
		}
		r.repo = repo
	}
}
