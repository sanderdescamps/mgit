package repo

import (
	"os"

	"github.com/go-git/go-git/v5"
)

type DisplayInterface interface {
	Error(msg string)
	Errorf(format string, a ...any)
	Warning(msg string)
	Warningf(format string, a ...any)
	Info(msg string)
	Infof(format string, a ...any)
	Debug(msg string)
	Debugf(format string, a ...any)
	Change(msg string)
	Changef(format string, a ...any)
	Ok(msg string)
	Okf(format string, a ...any)
	Skip(msg string)
	Skipf(format string, a ...any)
	Print(format string, a ...any)
	Writer() *os.File
}

type Repo struct {
	Url      string
	RepoPath string
	repo     *git.Repository
	Display  DisplayInterface
}
