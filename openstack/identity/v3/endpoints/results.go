package endpoints

import (
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a concrete
// Endpoint. An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*Endpoint, error) {
	var s struct {
		Endpoint *Endpoint `json:"endpoint"`
	}
	err := r.ExtractInto(&s)
	return s.Endpoint, err
}

func (r commonResult) ExtractDetail() (*EndpointDetail, error) {
	var s struct {
		EndpointDetail *EndpointDetail `json:"endpoint"`
	}
	err := r.ExtractInto(&s)
	return s.EndpointDetail, err
}

// CreateResult is the response from a Create operation. Call its Extract
// method to interpret it as an Endpoint.
type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as an Endpoint.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Endpoint describes the entry point for another service's API.
type Endpoint struct {
	// ID is the unique ID of the endpoint.
	ID string `json:"id"`

	// Availability is the interface type of the Endpoint (admin, internal,
	// or public), referenced by the gophercloud.Availability type.
	Availability gophercloud.Availability `json:"interface"`

	// Name is the name of the Endpoint.
	Name string `json:"name"`

	// Region is the region the Endpoint is located in.
	Region string `json:"region"`

	// ServiceID is the ID of the service the Endpoint refers to.
	ServiceID string `json:"service_id"`

	// URL is the url of the Endpoint.
	URL string `json:"url"`
}

type EndpointList struct {
	// ID is the unique ID of the endpoint.
	ID string `json:"id"`

	// Region is the region the Endpoint is located in.
	Region string `json:"region"`

	RegionID string `json:"region_id"`

	// ServiceID is the ID of the service the Endpoint refers to.
	ServiceID string `json:"service_id"`

	// Availability is the interface type of the Endpoint (admin, internal,
	// or public), referenced by the gophercloud.Availability type.
	Availability gophercloud.Availability `json:"interface"`

	// URL is the url of the Endpoint.
	URL string `json:"url"`

	Enabled bool `json:"enabled"`

	Links map[string]interface{} `json:"links"`
}

// Endpoint describes the entry point for another service's API.
type EndpointDetail struct {
	// ID is the unique ID of the endpoint.
	ID string `json:"id"`

	// Region is the region the Endpoint is located in.
	Region string `json:"region"`

	RegionID string `json:"region_id"`

	// ServiceID is the ID of the service the Endpoint refers to.
	ServiceID string `json:"service_id"`

	// Availability is the interface type of the Endpoint (admin, internal,
	// or public), referenced by the gophercloud.Availability type.
	Availability gophercloud.Availability `json:"interface"`

	// URL is the url of the Endpoint.
	URL string `json:"url"`

	Enabled bool `json:"enabled"`

	Links map[string]interface{} `json:"links"`
}

// EndpointPage is a single page of Endpoint results.
type EndpointPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if no Endpoints were returned.
func (r EndpointPage) IsEmpty() (bool, error) {
	es, err := ExtractEndpoints(r)
	return len(es) == 0, err
}

// ExtractEndpoints extracts an Endpoint slice from a Page.
func ExtractEndpoints(r pagination.Page) ([]Endpoint, error) {
	var s struct {
		Endpoints []Endpoint `json:"endpoints"`
	}
	err := (r.(EndpointPage)).ExtractInto(&s)
	return s.Endpoints, err
}

func ExtractEndpointsList(r pagination.Page) ([]EndpointList, error) {
	var s struct {
		Endpoints []EndpointList `json:"endpoints"`
	}
	err := (r.(EndpointPage)).ExtractInto(&s)
	return s.Endpoints, err
}
