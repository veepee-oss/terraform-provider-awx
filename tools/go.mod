module github.com/veepee-oss/terraform-provider-awx/tools

go 1.14

require (
	github.com/veepee-oss/terraform-provider-awx v1.0.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.10.1
	github.com/magefile/mage v1.11.0
	github.com/nolte/plumbing v0.0.1
)

replace github.com/veepee-oss/terraform-provider-awx => ../.
