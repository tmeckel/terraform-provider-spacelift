---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "spacelift_azure_devops_integration Data Source - terraform-provider-spacelift"
subcategory: ""
description: |-
  spacelift_azure_devops_integration returns details about Azure DevOps integration
---

# spacelift_azure_devops_integration (Data Source)

`spacelift_azure_devops_integration` returns details about Azure DevOps integration

## Example Usage

```terraform
data "spacelift_azure_devops_integration" "azure_devops_integration" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **organization_url** (String) Azure DevOps integration organization url
- **webhook_password** (String) Azure DevOps integration webhook password

