package projects

import (
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/pagination"
)

type projectResult struct {
	gophercloud.Result
}

// GetResult is the result of a Get request. Call its Extract method to
// interpret it as a Project.
type GetResult struct {
	projectResult
}

// CreateResult is the result of a Create request. Call its Extract method to
// interpret it as a Project.
type CreateResult struct {
	projectResult
}

// DeleteResult is the result of a Delete request. Call its ExtractErr method to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a Project.
type UpdateResult struct {
	projectResult
}

type Links struct {
	Next string `json:"next"`
	Previous string `json:"previous"`
	Self string `json:"self"`
}

// Project represents an OpenStack Identity Project.
type Project struct {
	// IsDomain indicates whether the project is a domain.
	IsDomain bool `json:"is_domain"`

	// Description is the description of the project.
	Description string `json:"description"`

	// DomainID is the domain ID the project belongs to.
	DomainID string `json:"domain_id"`

	// Enabled is whether or not the project is enabled.
	Enabled bool `json:"enabled"`

	// ID is the unique ID of the project.
	ID string `json:"id"`

	// Name is the name of the project.
	Name string `json:"name"`

	// ParentID is the parent_id of the project.
	ParentID string `json:"parent_id"`

	Links Links `json:"links"`
}

type Extra struct {

}


type ProjectUpdated struct {
	// IsDomain indicates whether the project is a domain.
	IsDomain bool `json:"is_domain"`

	// Description is the description of the project.
	Description string `json:"description"`

	// DomainID is the domain ID the project belongs to.
	DomainID string `json:"domain_id"`

	// Enabled is whether or not the project is enabled.
	Enabled bool `json:"enabled"`

	// ID is the unique ID of the project.
	ID string `json:"id"`

	// Name is the name of the project.
	Name string `json:"name"`

	// ParentID is the parent_id of the project.
	ParentID string `json:"parent_id"`

	Links Links `json:"links"`

	Extra Extra `json:"extra"`
}

// ProjectPage is a single page of Project results.
type ProjectPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Projects contains any results.
func (r ProjectPage) IsEmpty() (bool, error) {
	projects, err := ExtractProjects(r)
	return len(projects) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r ProjectPage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Links.Next, err
}

// ExtractProjects returns a slice of Projects contained in a single page of
// results.
func ExtractProjects(r pagination.Page) ([]Project, error) {
	var s struct {
		Projects []Project `json:"projects"`
	}
	err := (r.(ProjectPage)).ExtractInto(&s)
	return s.Projects, err
}

// Extract interprets any projectResults as a Project.
func (r projectResult) Extract() (*Project, error) {
	var s struct {
		Project *Project `json:"project"`
	}
	err := r.ExtractInto(&s)
	return s.Project, err
}

// Extract interprets any projectResults as a Project.
func (r projectResult) ExtractProjectUpdated() (*ProjectUpdated, error) {
	var s struct {
		Project *ProjectUpdated `json:"project"`
	}
	err := r.ExtractInto(&s)
	return s.Project, err
}

type ProlectList struct {
	Projects []Project `json:"projects"`
}
func (r projectResult) ExtractProjectList() ([]Project, error) {
	var s ProlectList
	err := r.ExtractInto(&s)
	return s.Projects, err
}