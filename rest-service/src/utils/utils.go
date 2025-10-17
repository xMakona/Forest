package utils

import (
	"path"
	"strings"
)

func ParsePathToken(url string) (token, rest string) {
	url = path.Clean("/" + url)
	idx := strings.Index(url[1:], "/") + 1
	if idx <= 0 {
		return url[1:], "/"
	}
	return url[1:idx], url[idx:]
}
