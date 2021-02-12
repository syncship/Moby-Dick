package commands

import (
	"regexp"

	"github.com/syncship/moby-dick/data/database"
	"github.com/syncship/moby-dick/pkg/router"
)

var db = database.New()

// MentionRegex ...
var MentionRegex = regexp.MustCompile("<(@!?\\d+)>")

// Router ...
var Router = router.New()
