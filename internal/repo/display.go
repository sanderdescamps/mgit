package repo

import (
	"io"

	"github.com/go-git/go-git/plumbing/protocol/packp/sideband"
)

type RepoDisplay struct {
	sideband.Progress
	writer io.Writer
}

func (d RepoDisplay) Write(p []byte) (n int, err error) {
	return d.writer.Write(p)
}
