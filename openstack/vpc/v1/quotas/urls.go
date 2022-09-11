package quotas

import (
	"github.com/DashuOps/gophercloud"
)

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
