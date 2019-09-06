package internal

import (
	"os"
)

// issueProvider is an interface for providing issue management using project management APIs (Jira, github issus, etc.)
// There is not yet a concept of a project management provider here in prwatch, but perhaps there will be. In the event
// that such a time arrives, this inteface will become the interface through which issue management is provided for
// project management providers.
type issueProvider interface {
	TransitionIssue(i issue) (ok bool)
	CommentIssue(i issue) (ok bool)
}

type issue struct {
	ID    string `json:"id,omitempty" structs:"id,omitempty"`
	Key   string `json:"key,omitempty" structs:"key,omitempty"`
	Value string `json:"value,omitempty" structs:"key,omitempty"`
	Owner string `json:"owner,omitempty" structs:"owner,omitempty"`
}

type issueComment struct {
	user    string
	comment string
}

// the issue status to transitions issues to upon conflict
var transitionName = os.Getenv("CONFLICT_ISSUE_STATUS")

func issueTransitionsEnabled() (yes bool) {

	// flag to enable issue transitions upon conflict
	if value, ok := os.LookupEnv("ENABLE_ISSUE_TRANSITIONS"); ok {
		yes = value == "true"
	}

	return
}

func issueCommentsEnabled() (yes bool) {

	// flag to enable issue comments upon conflict
	if value, ok := os.LookupEnv("ENABLE_ISSUE_COMMENTS"); ok {
		yes = value == "true"
	}

	return
}
