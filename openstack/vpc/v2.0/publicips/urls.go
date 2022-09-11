package publicips

import "github.com/DashuOps/gophercloud"

func CreateURL(c *gophercloud.ServiceClient)string{
	return c.ServiceURL("publicips")
}