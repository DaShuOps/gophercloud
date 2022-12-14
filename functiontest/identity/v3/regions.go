package main

import (
	"encoding/json"
	"fmt"
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/functiontest/common"
	"github.com/DashuOps/gophercloud/openstack"
	"github.com/DashuOps/gophercloud/openstack/identity/v3/regions"
)

func main() {

	fmt.Println("main start...")

	provider, err_auth := common.AuthAKSK()
	//provider, err_auth := common.AuthToken()

	if err_auth != nil {
		fmt.Println("Failed to get the provider: ", err_auth)
		return
	}

	sc, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{})

	if err != nil {
		fmt.Println("get IAM v3 failed")
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	GetRegion(sc)
	ListRegion(sc)

	fmt.Println("main end...")
}

// 查询区域详情
// Query the region detail
// GET /v3/regions/{region_id}
func GetRegion(client *gophercloud.ServiceClient) {
	fmt.Println("start TestGetRegion")
	result, err := regions.Get(client, "cn-north-7").Extract()

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(result, "", " ")
	fmt.Println(string(bytes))

	fmt.Println("finish TestGetRegion")
}

// 查询区域列表
// Query a region list
// GET /v3/regions
func ListRegion(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestListRegion")

	opts := regions.ListOpts{}

	resp, err := regions.List(sc, opts).AllPages()

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode", ue.ErrCode)
			fmt.Println("ErrMessage", ue.ErrMessage)
		}
		return
	}
	regionslist, err := regions.ExtractRegions(resp)

	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, d := range regionslist {

		b, _ := json.MarshalIndent(d, "", " ")
		fmt.Println(string(b))
	}

	fmt.Println("finish TestListRegion")
}
