package magentgo

import (
	"net/url"
	"regexp"
	"strings"
)

// basic check url is valid
func IsUrl(str string) bool {
	validProtocols := []string{
		"http",
		"https",
	}

	// https://stackoverflow.com/a/9284473
	pattern := `((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`

	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	// validate by regex
	if !re.MatchString(str) {
		return false
	}

	// validate using go url parser
	u, err := url.Parse(str)
	if err != nil {
		return false
	}

	// basic hostname checks
	if strings.HasPrefix(u.Hostname(), ".") || strings.HasSuffix(u.Hostname(), ".") || strings.Contains(u.Hostname(), "_") {
		return false
	}

	validCount := 0
	for _, protocol := range validProtocols {
		if u.Scheme == protocol {
			validCount++
		}
	}

	return validCount != 0
}
