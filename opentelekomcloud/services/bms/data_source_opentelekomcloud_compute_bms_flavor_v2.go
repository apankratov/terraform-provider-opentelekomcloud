package bms

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/bms/v2/flavors"

	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/common/cfg"
	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/services/ims"
)

func DataSourceBMSFlavorV2() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBMSFlavorV2Read,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"min_ram": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"vcpus": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"min_disk": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			"disk": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"swap": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"rx_tx_factor": {
				Type:     schema.TypeFloat,
				Computed: true,
			},

			"sort_key": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "id",
			},

			"sort_dir": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "asc",
				ValidateFunc: ims.DataSourceImagesImageV2SortDirection,
			},
		},
	}
}

func dataSourceBMSFlavorV2Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*cfg.Config)
	flavorClient, err := config.ComputeV2Client(config.GetRegion(d))
	if err != nil {
		return fmt.Errorf("Error creating OpenTelekom bms client: %s", err)
	}

	listOpts := flavors.ListOpts{
		MinDisk: d.Get("min_disk").(int),
		MinRAM:  d.Get("min_ram").(int),
		Name:    d.Get("name").(string),
		ID:      d.Id(),
		SortKey: d.Get("sort_key").(string),
		SortDir: d.Get("sort_dir").(string),
	}
	var flavor flavors.Flavor
	refinedflavors, err := flavors.List(flavorClient, listOpts)
	if err != nil {
		return fmt.Errorf("Unable to retrieve flavors: %s", err)
	}

	if len(refinedflavors) < 1 {
		return fmt.Errorf("Your query returned no results. " +
			"Please change your search criteria and try again.")
	} else {
		flavor = refinedflavors[0]
	}

	log.Printf("[DEBUG] Single Flavor found: %s", flavor.ID)
	d.SetId(flavor.ID)
	d.Set("name", flavor.Name)
	d.Set("disk", flavor.Disk)
	d.Set("min_disk", flavor.MinDisk)
	d.Set("min_ram", flavor.MinRAM)
	d.Set("ram", flavor.RAM)
	d.Set("rx_tx_factor", flavor.RxTxFactor)
	d.Set("swap", flavor.Swap)
	d.Set("vcpus", flavor.VCPUs)

	return nil
}