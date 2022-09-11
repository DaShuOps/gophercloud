package apiversions

import (
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/pagination"
)

// ListVersions lists all the Neutron API versions available to end-users
func ListVersions(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, apiVersionsURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}
