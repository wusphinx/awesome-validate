/*
规则参照:https://emailregex.com/
*/

package validate

import "regexp"

var emailRegexp *regexp.Regexp

func init() {
	str := `^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`
	emailRegexp = regexp.MustCompile(str)
}

func ValidateEmail(email string) bool {
	if !emailRegexp.MatchString(email) {
		return false
	}
	return true
}
