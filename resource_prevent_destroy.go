package main

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePreventDestroy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePreventDestroyCreate,
		Read:   resourcePreventDestroyRead,
		Delete: resourcePreventDestroyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"triggers": {
				Type:     schema.TypeMap,
				ForceNew: true,
			},
		},
	}
}

func resourcePreventDestroyCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(strconv.Itoa(rand.Int()))
	return resourcePreventDestroyRead(d, m)
}

func resourcePreventDestroyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePreventDestroyDelete(d *schema.ResourceData, m interface{}) error {
	if os.Getenv("TF_PREVENT_DESTROY") == "false" {
		log.Printf("TF_PREVENT_DESTROY=false, so allowing destroy.")
		return nil
	}
	return errors.New(
		"Destroy blocked by figma_prevent_destroy." +
			" This protects against accidental destruction of important resources." +
			" Please check your plan and make sure you're not destroying anything important." +
			" If you really must destroy, set TF_PREVENT_DESTROY=false and re-run.")
}
