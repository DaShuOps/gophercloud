package projects

import "github.com/DashuOps/gophercloud"


func updateURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

func getURL(client *gophercloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}