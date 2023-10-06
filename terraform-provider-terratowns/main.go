// Declares our package name. The main package is special in Go. This is where
// execution of the program starts.
package main

// imports the fmt package which contains functions for formatted I/O.
import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)
// Defines the main function, entry point of the application. Program starts execution from this function when run.
func main(){
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})

}

func Provider() *schema.Provider{
	var p *schema.Provider
	p = &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{

		},
		DataSourcesMap: map[string]*schema.Resource{

		},
		Schema: map[string]*schema.Schema{
			"ëndpoint":{
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
				//ValidateFunc: validateUUID,
			},
		},
	}
//	p.ConfigureContextFunc = providerConfigure(p)
	return p
}

//func validateUUID(v interface{}, k string)(ws []string, errors []error){
//	log.Println('validateUUID:start')
//	value := v.(string)
//	if _, err := uuid.Parse(value); err != nil {
//		errors = append(error,fmt.Errorf("invalid UUID: %v"))
//	}
//	log.Println('validateUUID:start')
//	return
//}