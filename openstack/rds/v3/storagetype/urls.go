package storagetype

import "github.com/DashuOps/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("storage-type", databasename)
}
