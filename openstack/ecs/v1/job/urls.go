package job

import "github.com/DashuOps/gophercloud"

// Querying task statuses URL
func jobURL(sc *gophercloud.ServiceClient, jobId string) string {
	return sc.ServiceURL("jobs", jobId)
}