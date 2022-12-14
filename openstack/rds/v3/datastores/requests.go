package datastores

import (
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/pagination"
)


func List(client *gophercloud.ServiceClient, databasesname string) pagination.Pager {
	url := listURL(client, databasesname)

	pageRdsList := pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DataStoresPage{pagination.Offset{PageResult: r}}
	})

	rdsheader := map[string]string{"Content-Type": "application/json"}
	pageRdsList.Headers = rdsheader
	return pageRdsList
}
