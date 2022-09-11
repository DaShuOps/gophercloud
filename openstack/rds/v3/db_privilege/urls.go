package db_privilege

import "github.com/DashuOps/gophercloud"

func createURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "db_privilege")
}

func deleteURL(sc *gophercloud.ServiceClient, instanceID string) string {
	return sc.ServiceURL("instances", instanceID, "db_privilege")
}
