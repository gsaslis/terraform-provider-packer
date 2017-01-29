package main

import (
    "log"

    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"work_dir": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PACKER_DIR", nil),
			},
			/*
			"builder_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PACKER_BUILDER", nil),
			},
			*/
			"json_path": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("PACKER_JSON_PATH", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"packer_build":  resourcePacker(),
		},

		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		WorkDir:  d.Get("work_dir").(string),
		//BuilderName: d.Get("builder_name").(string),
		JsonPath: d.Get("json_path").(string),
	}

	log.Println("[INFO] Initializing Packer Builder")
	return config.New()
}