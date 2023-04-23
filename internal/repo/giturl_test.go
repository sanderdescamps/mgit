package repo_test

import (
	"testing"

	"github.com/sanderdescamps/mgit/internal/repo"
)

func getTestUrls() []string {
	return []string{"ssh://user@host.xz:3022/path/to/repo.git/",
		"ssh://user@host.xz/path/to/repo.git/",
		"ssh://host.xz:3022/path/to/repo.git/",
		"ssh://host.xz/path/to/repo.git/",
		"ssh://user@host.xz/path/to/repo.git/",
		"ssh://host.xz/path/to/repo.git/",
		"ssh://user@host.xz/~user/path/to/repo.git/",
		"ssh://host.xz/~user/path/to/repo.git/",
		"ssh://user@host.xz/~/path/to/repo.git",
		"ssh://host.xz/~/path/to/repo.git",
		"user@host.xz:/path/to/repo.git/",
		"host.xz:/path/to/repo.git/",
		"user@host.xz:~user/path/to/repo.git/",
		"host.xz:~user/path/to/repo.git/",
		"user@host.xz:path/to/repo.git",
		"host.xz:path/to/repo.git",
		"rsync://host.xz/path/to/repo.git/",
		"git://host.xz/path/to/repo.git/",
		"git://host.xz/~user/path/to/repo.git/",
		"http://host.xz/path/to/repo.git/",
		"https://host.xz/path/to/repo.git/",
	}
}

func TestGitUrlParsing(t *testing.T) {
	for _, u := range getTestUrls() {

		parsedUrl, err := repo.ParseToGitUrl(u)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("parsed %s to %s", u, (*parsedUrl).GetUrl())
		}

	}
}
