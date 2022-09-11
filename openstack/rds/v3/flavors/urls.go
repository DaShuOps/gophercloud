package flavors

import "github.com/DashuOps/gophercloud"

func listURL(sc *gophercloud.ServiceClient, databasename string) string {
	return sc.ServiceURL("flavors", databasename)
}
