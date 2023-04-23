package repo_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/sanderdescamps/mgit/internal/repo"
)

func TestRepoUrlRegex(t *testing.T) {
	input := getTestUrls()

	regexLong := regexp.MustCompile(repo.GIT_URL_REGEX_LONG)
	regexShort := regexp.MustCompile(repo.GIT_URL_REGEX_SHORT)
	failed := []string{}
	for _, i := range input {
		if !regexLong.MatchString(i) && !regexShort.MatchString(i) {
			failed = append(failed, i)
		}
	}

	if len(failed) > 0 {
		t.Fatalf("Regex does not match following git urs's: %s", strings.Join(failed, ", "))
	}
}

// func TestTcpConnect(t *testing.T) {
// 	input := getTestUrls()

// 	for _, i := range input {
// 		pUrl, err := repo.ParseToGitUrl(i)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		repo.TcpConnectionTest(pUrl.GetHost(),)
// 	}

// 	if len(failed) > 0 {
// 		t.Fatalf("Regex does not match following git urs's: %s", strings.Join(failed, ", "))
// 	}
// }

func TestRepoCheck(t *testing.T) {
	for _, u := range getTestUrls() {

		parsedUrl, err := repo.ParseToGitUrl(u)
		if err != nil {
			t.Fatal(err)
		} else {
			t.Logf("parsed %s to %s", u, (*parsedUrl).GetUrl())
		}

	}
}
