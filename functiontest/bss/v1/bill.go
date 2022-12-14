package main

import (
	"encoding/json"
	"fmt"
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/functiontest/common"
	"github.com/DashuOps/gophercloud/openstack"
	"github.com/DashuOps/gophercloud/openstack/bss/v1/bill"
)

func main() {
 	fmt.Println("bill start...")

	//打开debug日志
	gophercloud.EnableDebug = true

	provider, err := common.AuthToken()
	//provider, err := common.AuthAKSK()
	if err != nil {
		fmt.Println("get provider client failed")
		fmt.Println(err.Error())
		return
	}
	fmt.Println("auth success!")

	// 初始化服务的client
	sc, err := openstack.NewBSSV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	TestQueryPartnerMonthlyBills(sc)
	TestQueryMonthlyExpenditureSummary(sc)
	TestQueryResourceUsageDetails(sc)
	TestQueryResourceUsageRecord(sc)
	fmt.Println("bill end...")
}

func TestQueryPartnerMonthlyBills(client *gophercloud.ServiceClient) {
	opts := bill.QueryPartnerMonthlyBillsOpts{
		CustomerId:           "061013372c00d3410f3fc017ee1e8ac0",
		Period:               "2018-08",
	}
	detailRsp,err := bill.QueryPartnerMonthlyBills(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(detailRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestPostPaidBill success")
}

func TestQueryMonthlyExpenditureSummary(client *gophercloud.ServiceClient) {
	opts := bill.QueryMonthlyExpenditureSummaryOpts{
		Cycle:                "2018-05",
	}
	detailRsp,err := bill.QueryMonthlyExpenditureSummary(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(detailRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestCustomerSum success")
}

func TestQueryResourceUsageDetails(client *gophercloud.ServiceClient) {
	opts := bill.QueryResourceUsageDetailsOpts{
		Cycle:                "2018-05",
	}
	detailRsp,err := bill.QueryResourceUsageDetails(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(detailRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestCustomerSum success")
}

func TestQueryResourceUsageRecord(client *gophercloud.ServiceClient) {
	opts := bill.QueryResourceUsageRecordOpts{
		StartTime:            "123",
		EndTime:              "",
		CloudServiceTypeCode: "",
		RegionCode:           "",
		OrderId:              "",
		PayMethod:            "",
		Offset:               0,
		Limit:                0,
		ResourceId:           "",
		EnterpriseProjectId:  "",
	}
	detailRsp,err := bill.QueryResourceUsageRecord(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(detailRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("TestCustomerSum success")
}