package availabilityzones

import "github.com/DashuOps/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}

func listDetailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-availability-zone", "detail")
}
