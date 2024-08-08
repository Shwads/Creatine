package scriptParser

import "strings"

func isFileName(name string) bool {
	return len(strings.Split(name, ".")) > 1
}
