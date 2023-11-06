package nhncloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/openstack/networking/v2/extensions/subnetpools"
)

func networkingSubnetpoolV2StateRefreshFunc(client *gophercloud.ServiceClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		subnetpool, err := subnetpools.Get(client, id).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return subnetpool, "DELETED", nil
			}
			if _, ok := err.(gophercloud.ErrDefault409); ok {
				return subnetpool, "ACTIVE", nil
			}

			return nil, "", err
		}

		return subnetpool, "ACTIVE", nil
	}
}