package roles

import (
	"encoding/json"
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/internal"
	"github.com/DashuOps/gophercloud/pagination"
)

// Role grants permissions to a user.
type Role struct {
	// DomainID is the domain ID the role belongs to.
	DomainID string `json:"domain_id"`

	// ID is the unique ID of the role.
	ID string `json:"id"`

	// Links contains referencing links to the role.
	Links map[string]interface{} `json:"links"`

	// Name is the role name
	Name string `json:"name"`

	// Extra is a collection of miscellaneous key/values.
	Extra map[string]interface{} `json:"-"`
}

type RoleDetail struct {
	// DomainID is the domain ID the role belongs to.
	DomainID string `json:"domain_id"`

	Catalog string `json:"catalog"`

	// ID is the unique ID of the role.
	ID string `json:"id"`

	// Links contains referencing links to the role.
	Links map[string]interface{} `json:"links"`

	// Name is the role name
	Name string `json:"name"`

	CreateTime string `json:"created_time"`

	Description string `json:"description"`

	DescriptionCn string `json:"description_cn"`

	DisplayName string `json:"display_name"`

	Flag string `json:"flag"`

	Type string `json:"type"`

	UpdatedTime string `json:"updated_time"`

	Policy map[string]interface{} `json:"policy"`
}

func (r *Role) UnmarshalJSON(b []byte) error {
	type tmp Role
	var s struct {
		tmp
		Extra map[string]interface{} `json:"extra"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Role(s.tmp)

	// Collect other fields and bundle them into Extra
	// but only if a field titled "extra" wasn't sent.
	if s.Extra != nil {
		r.Extra = s.Extra
	} else {
		var result interface{}
		err := json.Unmarshal(b, &result)
		if err != nil {
			return err
		}
		if resultMap, ok := result.(map[string]interface{}); ok {
			r.Extra = internal.RemainingKeys(Role{}, resultMap)
		}
	}

	return err
}

type roleResult struct {
	gophercloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Role.
type GetResult struct {
	roleResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Role
type CreateResult struct {
	roleResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a Role.
type UpdateResult struct {
	roleResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

type PutResult struct {
	gophercloud.ErrResult
}

type HeadResult struct {
	gophercloud.ErrResult
}

// RolePage is a single page of Role results.
type RolePage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Roles contains any results.
func (r RolePage) IsEmpty() (bool, error) {
	roles, err := ExtractRoles(r)
	return len(roles) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r RolePage) NextPageURL() (string, error) {
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

// ExtractProjects returns a slice of Roles contained in a single page of
// results.
func ExtractRoles(r pagination.Page) ([]Role, error) {
	var s struct {
		Roles []Role `json:"roles"`
	}
	err := (r.(RolePage)).ExtractInto(&s)
	return s.Roles, err
}

func ExtractListRoles(r pagination.Page) ([]RoleDetail, error) {
	var s struct {
		Roles []RoleDetail `json:"roles"`
	}
	err := (r.(RolePage)).ExtractInto(&s)
	return s.Roles, err
}

// Extract interprets any roleResults as a Role.
func (r roleResult) Extract() (*Role, error) {
	var s struct {
		Role *Role `json:"role"`
	}
	err := r.ExtractInto(&s)
	return s.Role, err
}

// Extract interprets any roleResults as a Role.
func (r roleResult) ExtractGroup() (*RoleDetail, error) {
	var s struct {
		Role *RoleDetail `json:"role"`
	}
	err := r.ExtractInto(&s)
	return s.Role, err
}

// RoleAssignment is the result of a role assignments query.
type RoleAssignment struct {
	Role  AssignedRole `json:"role,omitempty"`
	Scope Scope        `json:"scope,omitempty"`
	User  User         `json:"user,omitempty"`
	Group Group        `json:"group,omitempty"`
}

// AssignedRole represents a Role in an assignment.
type AssignedRole struct {
	ID string `json:"id,omitempty"`
}

// Scope represents a scope in a Role assignment.
type Scope struct {
	Domain  Domain  `json:"domain,omitempty"`
	Project Project `json:"project,omitempty"`
}

// Domain represents a domain in a role assignment scope.
type Domain struct {
	ID string `json:"id,omitempty"`
}

// Project represents a project in a role assignment scope.
type Project struct {
	ID string `json:"id,omitempty"`
}

// User represents a user in a role assignment scope.
type User struct {
	ID string `json:"id,omitempty"`
}

// Group represents a group in a role assignment scope.
type Group struct {
	ID string `json:"id,omitempty"`
}

// RoleAssignmentPage is a single page of RoleAssignments results.
type RoleAssignmentPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the RoleAssignmentPage contains no results.
func (r RoleAssignmentPage) IsEmpty() (bool, error) {
	roleAssignments, err := ExtractRoleAssignments(r)
	return len(roleAssignments) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to
// the next page of results.
func (r RoleAssignmentPage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next string `json:"next"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	return s.Links.Next, err
}

// ExtractRoleAssignments extracts a slice of RoleAssignments from a Collection
// acquired from List.
func ExtractRoleAssignments(r pagination.Page) ([]RoleAssignment, error) {
	var s struct {
		RoleAssignments []RoleAssignment `json:"role_assignments"`
	}
	err := (r.(RoleAssignmentPage)).ExtractInto(&s)
	return s.RoleAssignments, err
}

// AssignmentResult represents the result of an assign operation.
// Call ExtractErr method to determine if the request succeeded or failed.
type AssignmentResult struct {
	gophercloud.ErrResult
}

// UnassignmentResult represents the result of an unassign operation.
// Call ExtractErr method to determine if the request succeeded or failed.
type UnassignmentResult struct {
	gophercloud.ErrResult
}

type ListResult struct {
	roleResult
}

type ListResponse struct {
	Roles []struct {
		DomainID      string                 `json:"domain_id"`
		Catalog       string                 `json:"catalog"`
		ID            string                 `json:"id"`
		Links         map[string]interface{} `json:"links"`
		Name          string                 `json:"name"`
		CreateTime    string                 `json:"created_time"`
		Description   string                 `json:"description"`
		DescriptionCn string                 `json:"description_cn"`
		DisplayName   string                 `json:"display_name"`
		Flag          string                 `json:"flag"`
		Type          string                 `json:"type"`
		UpdatedTime   string                 `json:"updated_time"`
		Policy        map[string]interface{} `json:"policy"`
	} `json:"roles"`
	Links struct {
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
		Self     string      `json:"self"`
	} `json:"links"`
}

func (r ListResult) ExtractList() (*ListResponse, error) {
	var s ListResponse
	err := r.ExtractInto(&s)
	return &s, err
}
