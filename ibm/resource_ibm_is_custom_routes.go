package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

const ()

func resourceIBMCustomRoutes() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCustomRoutesCreate,
		Read:     resourceIBMCustomRoutesRead,
		Update:   resourceIBMCustomRoutesUpdate,
		Delete:   resourceIBMCustomRoutesDelete,
		Exists:   resourceIBMCustomRoutesExists,
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

			customRouteTableID: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			customRouteDestination: {
				Type:     schema.TypeString,
				Required: true,
			},

			customRouteZone: {
				Type:     schema.TypeString,
				Required: true,
			},

			customRouteAction: {
				Type:     schema.TypeString,
				Required: true,
			},

			customRouteName: {
				Type:     schema.TypeString,
				Required: true,
			},

			customRouteNextHop: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIBMCustomRoutesCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		log.Printf("error: %s", err)
		return err
	}

	var (
		vpcID   string
		tableID string
		//name    string
		action      string
		nextHop     string
		routeName   string
		destination string
		zone        string
	)

	vpcID = d.Get(customRouteTableVpcID).(string)
	tableID = d.Get(customRouteTableID).(string)
	//name = d.Get(customRouteTableName).(string)
	action = d.Get(customRouteAction).(string)
	nextHop = d.Get(customRouteNextHop).(string)
	routeName = d.Get(customRouteName).(string)
	destination = d.Get(customRouteDestination).(string)
	zone = d.Get(customRouteZone).(string)
	z := &vpcv1.ZoneIdentityByName{
		Name: core.StringPtr(zone),
	}

	nh := &vpcv1.RouteNextHopPrototypeRouteNextHopIP{
		Address: core.StringPtr(nextHop),
	}

	createVpcRoutingTableRouteOptions := sess.NewCreateVpcRoutingTableRouteOptions(vpcID, tableID, destination, z)
	createVpcRoutingTableRouteOptions.SetName(routeName)
	createVpcRoutingTableRouteOptions.SetZone(z)
	createVpcRoutingTableRouteOptions.SetDestination(destination)
	createVpcRoutingTableRouteOptions.SetNextHop(nh)
	createVpcRoutingTableRouteOptions.SetAction(action)

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
	//createVpcRoutingTableRouteOptions.SetName(name)
	//log.Printf("%+v\n", createVpcRoutingTableOptions.Routes[0].Zone)
	//log.Printf("VPC ID: %+s\n", *createVpcRoutingTableOptions.VpcID)
	//log.Printf("Name ID: %+s\n", *createVpcRoutingTableOptions.Name)
	//log.Printf("Route Action: %+s\n", *createVpcRoutingTableOptions.Routes[0].Action)
	//log.Printf("Route Name: %+s\n", *createVpcRoutingTableOptions.Routes[0].Name)
	//log.Printf("Route Destination: %+s\n", *createVpcRoutingTableOptions.Routes[0].Destination)
	response, _, err := sess.CreateVpcRoutingTableRoute(createVpcRoutingTableRouteOptions)
	if err != nil {
		log.Print("******************")
		log.Printf("CREATE CUSTOM ROUTE ERROR: %s", err)
		log.Print("******************")
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", vpcID, tableID, *response.ID))

	return resourceIBMCustomRoutesRead(d, meta)
}

func resourceIBMCustomRoutesRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableRouteOptions := sess.NewGetVpcRoutingTableRouteOptions(id_set[0], id_set[1], id_set[2])
	response, _, err := sess.GetVpcRoutingTableRoute(getVpcRoutingTableRouteOptions)
	if err != nil {
		return err
	}

	d.Set("id", response.ID)
	d.Set(customRouteTableID, response.ID)

	return nil
}

func resourceIBMCustomRoutesUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMPrivateDnsZoneRead(d, meta)
}

func resourceIBMCustomRoutesDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return err
	}

	id_set := strings.Split(d.Id(), "/")

	deleteVpcRoutingTableRouteOptions := sess.NewDeleteVpcRoutingTableRouteOptions(id_set[0], id_set[1], id_set[2])
	response, err := sess.DeleteVpcRoutingTableRoute(deleteVpcRoutingTableRouteOptions)
	if err != nil && response.StatusCode != 404 {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMCustomRoutesExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).VpcV1API()
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getVpcRoutingTableRouteOptions := sess.NewGetVpcRoutingTableRouteOptions(id_set[0], id_set[1], id_set[2])
	_, response, err := sess.GetVpcRoutingTableRoute(getVpcRoutingTableRouteOptions)
	if err != nil && response.StatusCode != 404 {
		if response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
