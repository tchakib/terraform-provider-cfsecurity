package cfsecurity

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAsg() *schema.Resource {

	return &schema.Resource{

		Read: dataSourceAsgRead,

		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAsgRead(d *schema.ResourceData, meta interface{}) error {
	manager := meta.(*Manager)

	err := refreshTokenIfExpired(manager)
	if err != nil {
		return err
	}

	secGroup, err := manager.client.GetSecGroupByName(d.Get("name").(string))
	if err != nil {
		return err
	}
	d.SetId(secGroup.GUID)
	return nil
}
