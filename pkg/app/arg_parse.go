package app

import (
	"fmt"
	"regexp"
)

// regular expression to match an app argument
var appRgx = regexp.MustCompile(`^(?:(?P<bucket>[a-zA-Z0-9-_.]+)/)?(?P<app>.*\.json$|[a-zA-Z0-9-_.]+)(?:@(?P<version>.*))?$`)

// parse an app argument
func ParseApp(app string) (ParsedApp, error) {
	matches := appRgx.FindStringSubmatch(app)
	p := ParsedApp{}
	if matches == nil || len(matches) <= 1 {
		return p, fmt.Errorf("Arg %v did not match valid app name syntax, must match `bucket/name@version` format", app)
	}
	p.Name = matches[2]
	p.Bucket = matches[1]
	p.Version = matches[3]
	return p, nil
}

type ParsedApp struct {
	Name    string
	Bucket  string
	Version string
}
