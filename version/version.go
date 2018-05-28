package version

import (
	"io/ioutil"
	"strings"
)

// GetVersion -
func GetVersion() (string, error) {
	v, err := ioutil.ReadFile("VERSION.txt")
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(v), "\n"), nil
}
