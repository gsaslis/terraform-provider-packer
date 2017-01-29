package main

import (
    //"fmt"
    "log"
    "github.com/hashicorp/terraform/helper/schema"
)

func resourcePacker() *schema.Resource {
    return &schema.Resource{
        Create: resourcePackerCreate,
        Read:   resourcePackerRead,
        //Update: resourcePackerUpdate,
        Delete: resourcePackerDelete,

        Schema: map[string]*schema.Schema{
	    /*
            "work_dir": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
                ForceNew: true,
            },
	    */
            "builder_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
                ForceNew: true,
            },
	    /*
            "path_to_json": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
                ForceNew: true,
            },
	    */
            "ami_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}

func resourcePackerCreate(data *schema.ResourceData, meta interface{}) error {
    	config := meta.(Config)

    	builder := data.Get("builder_name").(string)

    	log.Printf("[DEBUG] Will run packer build for builder: %#v, in dir: %s, for JSON: %s", builder, config.WorkDir, config.JsonPath)

        // TODO: implement
	    /*
    	do, err := client.DomainCreate(app, heroku.DomainCreateOpts{Hostname: hostname})
    	if err != nil {
    		return err
    	}
	    */

    	//data.SetId(do.ID)
    	//data.Set("hostname", do.Hostname)
    	//data.Set("cname", fmt.Sprintf("%s.herokuapp.com", app))
        data.Set("ami_id", "ami-123456")
        data.SetId("ami-123456")

    	log.Printf("[INFO] Packer Build for %#v generated AMI ID: %s", builder, data.Id())
    	return nil

}

func resourcePackerRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourcePackerUpdate(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourcePackerDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}