package extensions

import "github.com/DashuOps/gophercloud"

// ListExtensionURL generates the URL for the extensions resource collection.
func ListExtensionURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("extensions")
}
