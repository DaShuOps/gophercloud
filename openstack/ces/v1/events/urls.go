package events

import (
	"github.com/DashuOps/gophercloud"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("events")
}
