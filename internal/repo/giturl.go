package repo

import (
	"fmt"
	"regexp"
	"strconv"
)

type GitUrl struct {
	user     string
	host     string
	port     int
	protocol string
	path     string
}

func ParseToGitUrl(s string) (*GitUrl, error) {
	regexLong := regexp.MustCompile(GIT_URL_REGEX_LONG)
	regexShort := regexp.MustCompile(GIT_URL_REGEX_SHORT)
	if regexLong.MatchString(s) {
		parts := regexLong.FindStringSubmatch(s)
		var port int
		if parts[5] != "" {
			port, _ = strconv.Atoi(parts[5])
		}

		var protocol string
		if oneOf(parts[1], []string{"ssh", "http", "https", "rsync", "git"}) {
			protocol = parts[1]
		} else {
			return nil, fmt.Errorf("invalid protocol")
		}

		var user string
		if parts[3] != "" {
			user = parts[3]
		} else if parts[6] != "" {
			user = parts[6]
		}
		return &GitUrl{user: user, host: parts[4], port: port, protocol: protocol, path: parts[7]}, nil
	} else if regexShort.MatchString(s) {
		parts := regexShort.FindStringSubmatch(s)
		var port int
		if parts[4] != "" {
			port, _ = strconv.Atoi(parts[4])
		}

		var user string
		if parts[2] != "" {
			user = parts[2]
		}
		return &GitUrl{user: user, host: parts[3], port: port, protocol: "ssh", path: parts[6]}, nil
	}
	return nil, fmt.Errorf("invalid repo url format: %s", s)
}

func (r *GitUrl) GetEndpoint() string {
	result := r.host
	if r.port > 0 {
		result = fmt.Sprintf("%s:%d", result, r.port)
	}
	return result
}

func (r *GitUrl) GetRawEndpoint() string {
	result := r.host
	if r.port > 0 {
		result = fmt.Sprintf("%s:%d", result, r.port)
	} else if r.port == 0 && r.protocol == "https" {
		result = fmt.Sprintf("%s:%d", result, 443)
	} else if r.port == 0 && r.protocol == "http" {
		result = fmt.Sprintf("%s:%d", result, 80)
	} else if oneOf(r.protocol, []string{"ssh", "git"}) {
		result = fmt.Sprintf("%s:%d", result, 22)
	}
	return result

}

func (r *GitUrl) GetUrl() string {
	if oneOf(r.protocol, []string{"ssh", "git"}) && r.user != "" {
		return fmt.Sprintf("%s://%s%s", r.protocol, fmt.Sprintf("%s@%s", r.user, r.GetEndpoint()), r.path)
	} else {
		return fmt.Sprintf("%s://%s%s", r.protocol, r.GetEndpoint(), r.path)
	}
}
