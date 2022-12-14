package groups

import "github.com/DashuOps/gophercloud"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("groups")
}

func getURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("groups")
}

func updateURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func deleteURL(client *gophercloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func addUserToGroupURL(client *gophercloud.ServiceClient, groupID string, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}

func checkUserInGroupURL(client *gophercloud.ServiceClient, groupID string, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}

func removeUserFromGroupURL(client *gophercloud.ServiceClient, groupID string, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}
