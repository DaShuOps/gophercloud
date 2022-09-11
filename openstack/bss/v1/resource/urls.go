package resource

import "github.com/DashuOps/gophercloud"

func listURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources/detail")
}
