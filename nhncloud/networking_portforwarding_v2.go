package nhncloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	gophercloud "github.com/nhn/nhncloud.gophercloud"
	"github.com/nhn/nhncloud.gophercloud/openstack/networking/v2/extensions/layer3/portforwarding"
)

func networkingPortForwardingV2StateRefreshFunc(client *gophercloud.ServiceClient, fipID, pfID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		pf, err := portforwarding.Get(client, fipID, pfID).Extract()
		if err != nil {
			if _, ok := err.(gophercloud.ErrDefault404); ok {
				return pf, "DELETED", nil
			}

			return nil, "", err
		}

		return pf, "ACTIVE", nil
	}
}