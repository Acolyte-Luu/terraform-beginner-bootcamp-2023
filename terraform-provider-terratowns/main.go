// Declares our package name. The main package is special in Go. This is where
// execution of the program starts.
package main

// imports the fmt package which contains functions for formatted I/O.
import (

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

		log.Println("endpoint: "+ config.Endpoint)
		log.Println("token: "+ config.Token)
		log.Println("uuid: "+ config.UserUuid)

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

		Schema: map[string]*schema.Schema{
			"name":{
				Type: schema.TypeString,
				Required: true,
				Description: "name of home",
			},
			"description":{
				Type: schema.TypeString,
				Required: true,
				Description: "description of home",
			},
			"domain_name":{
				Type: schema.TypeString,
				Required: true,
				Description: "cloudfront domain name of home",
			},
			"town":{
				Type: schema.TypeString,
				Required: true,
				Description: "town which the home belongs to",
			},
			"content_version":{
				Type: schema.TypeInt,
				Required: true,
				Description: "the content version of the home",
			},
		},
	}
	log.Println("Resource:end")

	return resource
}

func resourceHouseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	log.Println("resourceHouseCreate:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	payload := map[string]interface{}{
		"name": d.Get("name").(string),
		"description": d.Get("description").(string),
		"domain_name": d.Get("domain_name").(string),
		"town": d.Get("town").(string),
		"content_version": d.Get("content_version").(int),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	// Construct the http request
	url := config.Endpoint+"/u/"+config.UserUuid+"/homes"
	log.Println("URL: "+ url)
	req, err := http.NewRequest("POST",url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}

	// Set headers
	req.Header.Set("Authorization","Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// parse response json
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil{
		return diag.FromErr(err)
	}

	// StatusOK = 200 http response code
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to create home resource, status_code: %d, status: %s, body: %s",resp.StatusCode,resp.Status, responseData))
	}

	// handle response status

	homeUUID := responseData["uuid"].(string)
	d.SetId(homeUUID)

	log.Println("resourceHouseCreate:end")

	return diags
}

func resourceHouseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("resourceHouseRead:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	homeUUID := d.Id()

	// construct the http request
	url := config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	log.Println("URL: "+ url)
	req, err := http.NewRequest("GET",url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set headers
	req.Header.Set("Authorization","Bearer "+config.Token)
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Accept","application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// StatusOK = 200 http response code
	var responseData map[string]interface{}
	if resp.StatusCode == http.StatusOK {
		// parse response json
		
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil{
			return diag.FromErr(err)
		}
		d.Set("name",responseData["name"].(string))
		d.Set("description",responseData["description"].(string))
		d.Set("domain_name",responseData["domain_name"].(string))
		d.Set("content_version",responseData["content_version"].(float64))
	}else if resp.StatusCode == http.StatusNotFound {
		d.SetId("")
	}else if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to read home resource, status_code: %d, status: %s, body: %s",resp.StatusCode,resp.Status,responseData))
	}


	log.Println("resourceHouseRead:end")

	return diags
}

func resourceHouseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("resourceHouseUpdate:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	homeUUID := d.Id()

	payload := map[string]interface{}{
		"name": d.Get("name").(string),
		"description": d.Get("description").(string),
		"content_version": d.Get("content_version").(int),
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	// construct the http request
	url := config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	log.Println("URL: "+ url)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}

	// Set headers
	req.Header.Set("Authorization","Bearer "+config.Token)
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Accept","application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// parse response json
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil{
		return diag.FromErr(err)
	}

	// StatusOK = 200 http response code
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to update home resource, status_code: %d, status: %s, body: %s",resp.StatusCode,resp.Status, responseData))
	}

	log.Println("resourceHouseUpdate:end")

	d.Set("name",payload["name"])
	d.Set("description",payload["description"])
	d.Set("content_version",payload["content_version"])

	return diags
}

func resourceHouseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("resourceHouseDelete:start")
	var diags diag.Diagnostics

	config := m.(*Config)

	homeUUID := d.Id()

	// construct the http request
	url := config.Endpoint+"/u/"+config.UserUuid+"/homes/"+homeUUID
	log.Println("URL: "+ url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}


	// Set headers
	req.Header.Set("Authorization","Bearer "+config.Token)
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("Accept","application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	// StatusOK = 200 http response code
	if resp.StatusCode != http.StatusOK {
		return diag.FromErr(fmt.Errorf("failed to delete home resource, status_code: %d, status: %s",resp.StatusCode,resp.Status))
	}

	d.SetId("")

	log.Println("resourceHouseDelete:end")

	return diags
}