/*
*TBD*

# Example Usage

```hcl

	resource "awx_workflow_job_template_node_credential" "baseconfig" {
	  workflow_job_template_node_id = awx_workflow_job_template_node.baseconfig.id
	  credential_id                 = awx_credential_machine.pi_connection.id
	}

```
*/
package awx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	awx "github.com/veepee-oss/goawx/client"
)

func resourceWorkflowJobTemplateNodeCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWorkflowJobTemplateNodeCredentialsCreate,
		DeleteContext: resourceWorkflowJobTemplateNodeCredentialsDelete,
		ReadContext:   resourceWorkflowJobTemplateNodeCredentialsRead,

		Schema: map[string]*schema.Schema{

			"workflow_job_template_node_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"credential_id": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWorkflowJobTemplateNodeCredentialsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.WorkflowJobTemplateNodeService
	workflowJobTemplateNodeID := d.Get("workflow_job_template_node_id").(int)
	_, err := awxService.GetWorkflowJobTemplateNodeByID(workflowJobTemplateNodeID, make(map[string]string))
	if err != nil {
		return buildDiagNotFoundFail("workflow job template node", workflowJobTemplateNodeID, err)
	}

	result, err := awxService.AssociateCredentials(workflowJobTemplateNodeID, map[string]interface{}{
		"id": d.Get("credential_id").(int),
	}, map[string]string{})

	if err != nil {
		return buildDiagnosticsMessage("Create: WorkflowJobTemplateNode not AssociateCredentials", "Fail to add credentials with Id %v, for Node ID %v, got error: %s", d.Get("credential_id").(int), workflowJobTemplateNodeID, err.Error())
	}

	d.SetId(strconv.Itoa(result.ID))
	return diags
}

func resourceWorkflowJobTemplateNodeCredentialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceWorkflowJobTemplateNodeCredentialsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*awx.AWX)
	awxService := client.WorkflowJobTemplateNodeService
	workflowJobTemplateNodeID := d.Get("workflow_job_template_node_id").(int)
	res, err := awxService.GetWorkflowJobTemplateNodeByID(workflowJobTemplateNodeID, make(map[string]string))
	if err != nil {
		return buildDiagNotFoundFail("workflow job template node", workflowJobTemplateNodeID, err)
	}

	_, err = awxService.DisAssociateCredentials(res.ID, map[string]interface{}{
		"id": d.Get("credential_id").(int),
	}, map[string]string{})
	if err != nil {
		return buildDiagDeleteFail("WorkflowJobTemplateNode DisAssociateCredentials", fmt.Sprintf("DisAssociateCredentials %v, from Node ID %v got %s ", d.Get("credential_id").(int), d.Get("workflow_job_template_node_id").(int), err.Error()))
	}

	d.SetId("")
	return diags
}
