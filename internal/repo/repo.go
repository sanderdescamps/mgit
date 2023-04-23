package repo

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/sanderdescamps/mgit/internal/console"
)

// type Repo struct {
// 	Url      string
// 	RepoPath string
// 	repo     *git.Repository
// 	Display  DisplayInterface
// }

func NewRepo(url string, path string, display *console.Display, insecure bool) Repo {
	repo := Repo{
		Url:      url,
		RepoPath: path,
		Display:  *display,
		Insecure: false,
	}
	return repo
}

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

func (r *Repo) IsValidRemote() bool {
	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{r.Url},
	})
	_, err := rem.List(&git.ListOptions{})
	if err == nil {
		return true
	} else if err.Error() == "remote repository is empty" {
		return true
	}
	return false
}

func (r *Repo) Clone() {
	if r.repo != nil {
		r.Display.Info("Repo already cloned")
	} else if !r.isClonedOnFileSystem() {
		r.Display.Debugf("Clone repo %s...", r.Url)
		d := r.Display
		repo, err := git.PlainClone(r.RepoPath, false, &git.CloneOptions{
			URL:      r.Url,
			Progress: d.InfoWriter(),
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
			r.Display.Infof("Repo already cloned: %s", r.Url)
		} else {
			r.Display.Errorf("failed to load repo: %s\n%v", r.Url, err.Error())
			return
		}
		r.repo = repo
	}
}

func (r *Repo) LocalHead() string {
	ref, err := r.repo.Head()
	if err != nil {
		return ""
	} else {
		return ref.Hash().String()
	}
}

// func (r *Repo) RemoteHead() string {
// 	remote, err := r.repo.Remote("origin")
// 	if err != nil {
// 		return ""
// 	} else {
// 		return ref.Hash().String()
// 	}
// }
