package availabilityzones

import "github.com/DashuOps/gophercloud"

// listURL generates URL for list avaliabilityzones
func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
