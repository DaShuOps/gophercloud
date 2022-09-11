package configurations

import "github.com/DashuOps/gophercloud"

func listURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("configurations")
}

func createURL(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("configurations")
}
