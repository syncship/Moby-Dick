package commands

import (
	"regexp"

	"github.com/syncship/moby-dick/pkg/router"
)

// MentionRegex ...
var MentionRegex = regexp.MustCompile(`<(@!?\d+)>`)

// Router ...
var Router = router.New()
