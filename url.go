package validate

import (
	"net/url"
	"regexp"
)

var urlRegexp *regexp.Regexp

func init() {
	urlRegexp = regexp.MustCompile(`^(https)|(http)://[^/].*`)
}

func ValidateUrl(rawurl string) bool {
	if ok := urlRegexp.MatchString(rawurl); !ok {
		return false
	}

	if _, err := url.Parse(rawurl); err != nil {
		return false
	}

	return true
}
