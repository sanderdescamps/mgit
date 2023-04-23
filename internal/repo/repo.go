package repo

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/sanderdescamps/mgit/internal/console"
)

const (
	GIT_URL_REGEX_LONG  = `^(?:(\w+):\/\/)((?:(\w+)@)?((?:\w+\.)+\w+)(?::(\d*))?)(?:\/~(\w+))?([\/:][~\w\-\/\.]+)$`
	GIT_URL_REGEX_SHORT = `^((?:(\w+)@)?((?:\w+\.)+\w+)(?::(\d*))?)(?:\/~(\w+))?([\/:][~\w\-\/\.]+)$`
)

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

	_, err := rem.List(&git.ListOptions{InsecureSkipTLS: r.Insecure})
	if err == nil {
		return true
	} else if err.Error() == "remote repository is empty" {
		return true
	}
	return false
}

func (r *Repo) Clone() error {
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
			return err
		}
		r.repo = repo
	} else {
		repo, err := git.PlainOpen(r.RepoPath)
		if err == nil {
			r.Display.Infof("Repo already cloned: %s", r.Url)
		} else {
			r.Display.Errorf("failed to load repo: %s\n%v", r.Url, err.Error())
			return err
		}
		r.repo = repo
	}
	return nil
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

func (r *Repo) commitHashes() []string {
	result := []string{}
	commIter, err := r.repo.Log(&git.LogOptions{All: true})
	if err != nil {
		r.Display.Error("Failed to get log from repo")
	}
	for true {
		c, err := commIter.Next()
		if err != nil {
			r.Display.Error("failed to get git log")
		} else if c != nil {
			result = append(result, c.Hash.String())
		} else {
			break
		}
	}
	return result
}

func (r *Repo) branchNames() []string {
	result := []string{}
	branchIter, err := r.repo.Branches()
	if err != nil {
		r.Display.Error("Failed to get log from repo")
	}
	for true {
		b, err := branchIter.Next()
		if err != nil {
			r.Display.Error("failed to get git branches")
		} else if b != nil {
			result = append(result, b.Name().String())
		} else {
			break
		}
	}
	return result
}

func (r *Repo) PullBranch(branch string) error {
	wt, err := r.repo.Worktree()
	if err != nil {
		r.Display.Error("Failed to get work tree")
	}

	currentBranch := ""
	if head, err := r.repo.Head(); err == nil && head.Name().IsBranch() {
		currentBranch = head.Name().String()
	} else if err != nil {
		if !head.Name().IsBranch() {
			return errors.New(fmt.Sprintf("Failed to pull %s branch because repo has deatached head", branch))
		} else {
			return err
		}
	}

	if err := wt.Checkout(&git.CheckoutOptions{Branch: plumbing.ReferenceName(branch)}); err != nil {
		r.Display.Errorf("failed to checkout %s branch", branch)
	}
	if err := wt.Pull(&git.PullOptions{
		Progress: RepoDisplay{writer: r.Display.InfoWriter()},
	}); err != nil {
		r.Display.Errorf("failed to pull changes for %s branch", branch)
	}

	if err := wt.Checkout(&git.CheckoutOptions{Branch: plumbing.ReferenceName(currentBranch)}); err != nil {
		r.Display.Errorf("failed to restore repo to %s branch", currentBranch)
	}
	return nil
}

func (r *Repo) PullAllBranch(branch string) error {
	wt, err := r.repo.Worktree()
	if err != nil {
		r.Display.Error("Failed to get work tree")
	}

	currentBranch := ""
	if head, err := r.repo.Head(); err == nil && head.Name().IsBranch() {
		currentBranch = head.Name().String()
	} else if err != nil {
		if !head.Name().IsBranch() {
			return fmt.Errorf("failed to pull %s branch because repo has deatached head", branch)
		} else {
			return err
		}
	}

	branchList := r.branchNames()
	for _, b := range branchList {
		if err := wt.Checkout(&git.CheckoutOptions{Branch: plumbing.ReferenceName(b)}); err != nil {
			r.Display.Errorf("failed to checkout %s", b)
		}
		if err := wt.Pull(&git.PullOptions{
			Progress: RepoDisplay{writer: r.Display.InfoWriter()},
		}); err != nil {
			r.Display.Errorf("failed to pull changes for %s branch", b)
		}
	}

	if err := wt.Checkout(&git.CheckoutOptions{Branch: plumbing.ReferenceName(currentBranch)}); err != nil {
		r.Display.Errorf("failed to restore repo to %s branch", currentBranch)
	}
	return nil
}

func (r *Repo) CheckRepo(branch string) error {
	pUrl, err := ParseToGitUrl(r.Url)
	if err != nil {
		return err
	}
	_, err = net.LookupIP(pUrl.host)
	if err != nil {
		return fmt.Errorf("failed to resolve %s", pUrl.host)
	}

	err = tcpConnectionTest(pUrl.host, pUrl.port)
	if err != nil {
		return fmt.Errorf("failed to connect to %s", pUrl.host)
	}
	return nil
}

func (r *Repo) CheckTcpConnect() error {
	pUrl, err := ParseToGitUrl(r.Url)
	if err != nil {
		return err
	}

	//DNS resolve
	_, err = net.LookupIP(pUrl.host)
	if err != nil {
		return fmt.Errorf("failed to resolve %s", pUrl.host)
	}

	endpoint := pUrl.GetRawEndpoint()
	r.Display.Debugf("test endpoint: %s", endpoint)
	conn, _ := net.Dial("tcp", endpoint)

	err = conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		return fmt.Errorf("failed to set KeepAlive")
	}

	err = conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		return fmt.Errorf("failed to set KeepAlive period")
	}
	notify := make(chan error)

	go func() {
		buf := make([]byte, 1024)
		defer conn.Close()
		for {
			_, err := conn.Read(buf)
			if err != nil {
				notify <- err
				if io.EOF == err {
					fmt.Printf("EOF: %s", err)
					return
				} else {
					fmt.Printf("unexpected error: %s", err)
				}
			} else {
				notify <- nil
				return
			}
		}
	}()

	for {
		select {
		case err := <-notify:
			if err != nil {
				fmt.Println("connection dropped message", err)
				return err
			} else {
				return nil
			}
		case <-time.After(time.Second * 3):
			return fmt.Errorf("timeout connection")
		}
	}
}
