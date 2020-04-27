package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	customRouteTableVpcID          = "vpc_id"
	customRouteTableName           = "name"
	customRouteAction              = "action"
	customRouteDestination         = "destination"
	customRouteName                = "route_name"
	customRouteNextHop             = "next_hop"
	customRouteZone                = "zone"
	customRouteTableID             = "route_table_id"
	customRouteTableHref           = "href"
	customRouteTableCreatedAt      = "created_at"
	customRouteTableResourceType   = "resource_type"
	customRouteTableLifecycleState = "state"
)

func resourceIBMISRouteTable() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMISRouteTableCreate,
		Read:     resourceIBMISRouteTableRead,
		Update:   resourceIBMISRouteTableUpdate,
		Delete:   resourceIBMISRouteTableDelete,
		Exists:   resourceIBMISRouteTableExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			customRouteTableVpcID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			customRouteTableName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			customRouteAction: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			customRouteDestination: {
				Type:     schema.TypeString,
				Required: true,
			},

			customRouteName: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			customRouteNextHop: {
				Type:     schema.TypeString,
				Required: true,
			},

			customRouteZone: {
				Type:     schema.TypeString,
				Required: true,
			},

			customRouteTableID: {
				Type:     schema.TypeString,
				Computed: true,
			},

			customRouteTableHref: {
				Type:     schema.TypeString,
				Computed: true,
			},

			customRouteTableCreatedAt: {
				Type:     schema.TypeString,
				Computed: true,
			},

			customRouteTableResourceType: {
				Type:     schema.TypeString,
				Computed: true,
			},

			customRouteTableLifecycleState: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMISRouteTableCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		log.Printf("error: %s", err)
		return err
	}

	var (
		vpcID string
		name  string
		//action      string
		//nextHop     string
		//routeName   string
		//destination string
		//zone        string
	)

	vpcID = d.Get(customRouteTableVpcID).(string)
	name = d.Get(customRouteTableName).(string)
	//action = d.Get(customRouteAction).(string)
	//nextHop = d.Get(customRouteNextHop).(string)
	//routeName = d.Get(customRouteName).(string)
	//destination = d.Get(customRouteDestination).(string)
	//zone = d.Get(customRouteZone).(string)

	createVpcRoutingTableOptions := sess.NewCreateVpcRoutingTableOptions(vpcID)

	/* Construct an instance of the RoutePrototype model
	routePrototypeModel := []vpcv1.RoutePrototype{
		{
			Action: core.StringPtr(action),
			NextHop: &vpcv1.RouteNextHopPrototypeRouteNextHopIP{
				Address: core.StringPtr(nextHop),
			},
			Name:        core.StringPtr(routeName),
			Destination: core.StringPtr(destination),
			Zone: &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr(zone),
			},
		},
	}
	*/

	//createVpcRoutingTableOptions.SetRoutes(routePrototypeModel)
	//createVpcRoutingTableOptions.SetVpcID(vpcID)
	createVpcRoutingTableOptions.SetName(name)
	//log.Printf("%+v\n", createVpcRoutingTableOptions.Routes[0].Zone)
	//log.Printf("VPC ID: %+s\n", *createVpcRoutingTableOptions.VpcID)
	//log.Printf("Name ID: %+s\n", *createVpcRoutingTableOptions.Name)
	//log.Printf("Route Action: %+s\n", *createVpcRoutingTableOptions.Routes[0].Action)
	//log.Printf("Route Name: %+s\n", *createVpcRoutingTableOptions.Routes[0].Name)
	//log.Printf("Route Destination: %+s\n", *createVpcRoutingTableOptions.Routes[0].Destination)
	response, _, err := sess.CreateVpcRoutingTable(createVpcRoutingTableOptions)
	if err != nil {
		log.Print("******************")
		log.Printf("CREATE TABLE ERROR: %s", err)
		log.Print("******************")
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", vpcID, *response.ID))

	return resourceIBMISRouteTableRead(d, meta)
}

func resourceIBMISRouteTableRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableOptions := sess.NewGetVpcRoutingTableOptions(id_set[0], id_set[1])
	response, _, err := sess.GetVpcRoutingTable(getVpcRoutingTableOptions)
	if err != nil {
		return err
	}

	d.Set("id", response.ID)
	d.Set(customRouteTableID, response.ID)
	d.Set(customRouteTableHref, response.Href)
	d.Set(customRouteTableName, response.Name)
	d.Set(customRouteTableResourceType, response.ResourceType)
	d.Set(customRouteTableCreatedAt, response.CreatedAt)
	d.Set(customRouteTableLifecycleState, response.LifecycleState)

	return nil
}

func resourceIBMISRouteTableUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMPrivateDnsZoneRead(d, meta)
}

func resourceIBMISRouteTableDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")

	deleteZoneOptions := sess.NewDeleteVpcRoutingTableOptions(id_set[0], id_set[1])
	response, err := sess.DeleteVpcRoutingTable(deleteZoneOptions)
	if err != nil && response.StatusCode != 404 {
		log.Printf("Error deleting dns zone:%s", response)
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMISRouteTableExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableOptions := sess.NewGetVpcRoutingTableOptions(id_set[0], id_set[1])
	_, response, err := sess.GetVpcRoutingTable(getVpcRoutingTableOptions)
	if err != nil && response.StatusCode != 404 {
		if response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
