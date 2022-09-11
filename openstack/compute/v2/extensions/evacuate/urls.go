package evacuate

import (
	"github.com/DashuOps/gophercloud"
)

func actionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
