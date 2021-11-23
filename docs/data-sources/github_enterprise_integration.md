---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "spacelift_github_enterprise_integration Data Source - terraform-provider-spacelift"
subcategory: ""
description: |-
  spacelift_github_enterprise_integration returns details about Github Enterprise integration
---

# spacelift_github_enterprise_integration (Data Source)

`spacelift_github_enterprise_integration` returns details about Github Enterprise integration

## Example Usage

```terraform
data "spacelift_github_enterprise_integration" "github_enterprise_integration" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **api_host** (String) Github integration api host
- **app_id** (String) Github integration app id
- **webhook_secret** (String) Github integration webhook secret

