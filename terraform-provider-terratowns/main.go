// Declares our package name. The main package is special in Go. This is where
// execution of the program starts.
package main

// imports the fmt package which contains functions for formatted I/O.
import (
	"log"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)
// Defines the main function, entry point of the application. Program starts execution from this function when run.
func main(){
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})

}

type Config struct {
	Endpoint string
	Token string
	UserUuid string
}

func Provider() *schema.Provider{
	var p *schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"terratowns_home": Resource(),

		},
		DataSourcesMap: map[string]*schema.Resource{

		},
		Schema: map[string]*schema.Schema{
			"endpoint":{

				Type: schema.TypeString,
				Required: true,
				Description: "The endpoint for the external service",

			},
			"token":{
				Type: schema.TypeString,
				Required: true,
				Sensitive: true,
				Description: "Bearer token for authorization",

			},
			"user_uuid":{
				Type: schema.TypeString,
				Required: true,
				Description: "UUID for configuration",
				ValidateFunc: validateUUID,
			},
		},
	}
	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

func validateUUID(v interface{}, k string)(ws []string, errors []error){
	log.Println("validateUUID:start")
	value := v.(string)
	if _, err := uuid.Parse(value); err != nil {
		errors = append(errors,fmt.Errorf("invalid UUID format"))
	}
	log.Println("validateUUID:start")
	return
}

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc{
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics){
		log.Println("providerConfigure:start")
		config := Config{
			Endpoint: d.Get("endpoint").(string),
			Token: d.Get("token").(string),
			UserUuid: d.Get("user_uuid").(string),
		}
		log.Println("providerConfigure:end")
		return &config, nil
	}
}

func Resource() *schema.Resource{
	log.Println("Resource:start")
	resource := &schema.Resource{
		CreateContext: resourceHouseCreate,
		ReadContext: resourceHouseRead,
		UpdateContext: resourceHouseUpdate,
		DeleteContext: resourceHouseDelete,
	}
	log.Println("Resource:start")
	return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}