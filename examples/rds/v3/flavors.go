package main

import (
	"encoding/json"
	"fmt"
	"github.com/DashuOps/gophercloud"
	//"github.com/DashuOps/gophercloud/auth/token"
	"github.com/DashuOps/gophercloud/auth/aksk"
	"github.com/DashuOps/gophercloud/openstack"
	"github.com/DashuOps/gophercloud/openstack/rds/v3/flavors"
	"github.com/DashuOps/gophercloud/pagination"
	
)


func main() {

	fmt.Println("rds flavors test  start...")
	gophercloud.EnableDebug = true
	//Set authentication parameters
	akskopts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.xxx.yyy.com/v3",
		ProjectID:        "{ProjectID}",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
		Cloud:            "yyy.com",
		Region:           "xxx",
		DomainID:         "{domainID}",
	}
	provider, authErr := openstack.AuthenticatedClient(akskopts)
	if authErr != nil {
		fmt.Println("Failed to get the AuthenticatedClient: ", authErr)
		fmt.Println("Failed to get the provider: ", provider)
		return
	}

	client, clientErr := openstack.NewRDSV3(provider, gophercloud.EndpointOpts{Region:"xxx"})
	if clientErr != nil {
		fmt.Println("Failed to get the NewRDSV3 client: ", clientErr)
		return
	}

	ListDbFlavorsTest(client, "MySQL")
}

func ListDbFlavorsTest(sc *gophercloud.ServiceClient, databasename string) {
	opts := flavors.DbFlavorsOpts{
		Versionname: "5.7",
	}
	err := flavors.List(sc, opts, databasename).EachPage(func(page pagination.Page) (bool, error) {
		resp, pageErr := flavors.ExtractDbFlavors(page)
		if pageErr != nil {
			fmt.Println(pageErr)
			if ue, ok := pageErr.(*gophercloud.UnifiedError); ok {
				fmt.Println("ErrCode:", ue.ErrorCode())
				fmt.Println("Message:", ue.Message())
			}
			return false, pageErr
		}
		for _, v := range resp.Flavorslist {
			jsServer, _ := json.MarshalIndent(v, "", "   ")
			fmt.Println("Flavors log info is :", string(jsServer))

		}
		// When returns false, current page of data will be returned.
		// Otherwise,when true,all pages of data will be returned.
		return true, nil
	})

	if err != nil {
		fmt.Println(err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
}
