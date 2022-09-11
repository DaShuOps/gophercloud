package metrics

import "github.com/DashuOps/gophercloud"

func getMetricsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}
