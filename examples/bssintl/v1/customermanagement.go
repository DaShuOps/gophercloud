package main

import (
	"encoding/json"
	"fmt"
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/auth/aksk"
	"github.com/DashuOps/gophercloud/openstack"
	"github.com/DashuOps/gophercloud/openstack/bssintl/v1/customermanagement"
)

func main() {
	//AKSK auth，initial parameter.
	opts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		DomainID:         "{domainID}",
	}
	//initial provider client。
	provider, errAuth := openstack.AuthenticatedClient(opts)
	if errAuth != nil {
		fmt.Println("get provider client failed")
		fmt.Println(errAuth.Error())
		return
	}

	// initial client
	sc, err := openstack.NewBSSIntlV1(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Println("get bss client failed")
		fmt.Println(err.Error())
		return
	}

	CheckCustomerRegisterInfo(sc)

	CreateCustomer(sc)

	QueryCustomer(sc)

	FrozenCustomer(sc)

	UnFrozenCustomer(sc)

}

func CheckCustomerRegisterInfo(client *gophercloud.ServiceClient) {
	opts := customermanagement.CheckCustomerRegisterInfoOpts{
		SearchType: "name",
		SearchKey: "bss 02",
	}

	checkUserRsp,err := customermanagement.CheckCustomerRegisterInfo(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	if checkUserRsp.ErrorCode == "CBC.0000" {
		fmt.Println("CheckUserName success, status=", checkUserRsp.Status)
	}else {
		fmt.Println("CheckUserName failed, ErrorCode=", checkUserRsp.ErrorCode, checkUserRsp.ErrorMsg)
	}
}

func CreateCustomer(client *gophercloud.ServiceClient) {
	opts := customermanagement.CreateCustomerOpts{
		DomainName:       "xxxxxxx",
		XAccountId:       "xxxxxxx",
		XAccountType:     "xxxxxxx",
		Password:         "xxxxxxx",
	}

	createCustomerRsp,err := customermanagement.CreateCustomer(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}


	bytes, _ := json.MarshalIndent(createCustomerRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("CreateCustomer success")
}

func QueryCustomer(client *gophercloud.ServiceClient) {
	opts := customermanagement.QueryCustomerOpts{
		DomainName:           "xxxxxxx",
	}

	queryCustomerRsp,err := customermanagement.QueryCustomer(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(queryCustomerRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("QueryCustomer success")
}

func FrozenCustomer(client *gophercloud.ServiceClient) {
	opts := customermanagement.FrozenCustomerOpts{
		CustomerIds: []string{"xxxxxxx"},
		Reason: "abc",
	}

	queryCustomerRsp,err := customermanagement.FrozenCustomer(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(queryCustomerRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("QueryCustomer success")
}

func UnFrozenCustomer(client *gophercloud.ServiceClient) {
	opts := customermanagement.UnFrozenCustomerOpts{
		CustomerIds: []string{"xxxxxxx"},
		Reason: "abc",
	}

	queryCustomerRsp,err := customermanagement.UnFrozenCustomer(client, opts).Extract()

	if err != nil {
		fmt.Println("err:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	bytes, _ := json.MarshalIndent(queryCustomerRsp, "", " ")
	fmt.Println(string(bytes))
	fmt.Println("QueryCustomer success")
}