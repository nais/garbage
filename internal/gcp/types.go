package gcp

// What do you mean "it's not a string"?!

import (
	"fmt"
)

type ProjectLocationID struct {
	Project  string
	Location string
}

func (p ProjectLocationID) String() string {
	return fmt.Sprintf("projects/%s/locations/%s", p.Project, p.Location)
}

type RepositoryID struct {
	Project    string
	Location   string
	Repository string
}

func (p RepositoryID) String() string {
	return fmt.Sprintf("projects/%s/locations/%s/repositories/%s", p.Project, p.Location, p.Repository)
}
