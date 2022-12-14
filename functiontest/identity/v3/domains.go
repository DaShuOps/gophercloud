package main

import (
	"encoding/json"
	"fmt"
	"github.com/DashuOps/gophercloud"
	"github.com/DashuOps/gophercloud/functiontest/common"
	"github.com/DashuOps/gophercloud/openstack"
	"github.com/DashuOps/gophercloud/openstack/identity/v3/domains"
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
	GetAccessibleDomainDetail(sc)
	GetDomainPasswordStrengthPolicy(sc)
	GetDomainPasswordStrengthPolicyWithOption(sc)

	fmt.Println("main end...")
}

// 查询用户可以访问的账号详情
// Query details of account that can be accessed by users
// GET /v3/auth/domains
func GetAccessibleDomainDetail(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestGetAccessibleDomainDetail")
	resp := domains.ListDomains(sc)
	result, err := resp.ExtractList()
	if err != nil {
		fmt.Println("get domain list failed")
		return
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}

	fmt.Println("finish TestGetAccessibleDomainDetail")
}

// 查询账号密码强度策略
// Query the password strength policy of the account
// GET /v3/domains/{domain_id}/config/security_compliance
func GetDomainPasswordStrengthPolicy(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestGetDomainPasswordStrengthPolicy")
	domainID := ""
	resp := domains.GetDoaminPwdStrengthPolicy(sc, domainID)
	result, err := resp.ExtractPwdStrengthPolicy()
	if err != nil {
		fmt.Println("get domain password strength policy failed")
		return
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestGetDomainPasswordStrengthPolicy")
}

// 按条件查询账号密码强度策略
// Query the password strength policy of the account by condition
// GET /v3/domains/{domain_id}/config/security_compliance/{option}
func GetDomainPasswordStrengthPolicyWithOption(sc *gophercloud.ServiceClient) {
	fmt.Println("start TestGetDomainPasswordStrengthPolicyWithOption")
	domainID := ""
	option := "password_regex_description"
	resp := domains.GetDoaminPwdStrengthPolicyByOption(sc, domainID, option)
	result, err := resp.ExtractPwdStrengthPolicy()
	if err != nil {
		fmt.Println("get domain password strength policy failed")
		return
	} else {
		bytes, _ := json.MarshalIndent(result, "", " ")
		fmt.Println(string(bytes))
	}
	fmt.Println("finish TestGetDomainPasswordStrengthPolicyWithOption")
}
