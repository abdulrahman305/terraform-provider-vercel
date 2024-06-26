---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "vercel_edge_config_schema Data Source - terraform-provider-vercel"
subcategory: ""
description: |-
  An Edge Config Schema provides an existing Edge Config with a JSON schema. Use schema protection to prevent unexpected updates that may cause bugs or downtime.
---

# vercel_edge_config_schema (Data Source)

An Edge Config Schema provides an existing Edge Config with a JSON schema. Use schema protection to prevent unexpected updates that may cause bugs or downtime.

## Example Usage

```terraform
data "vercel_edge_config_schema" "test" {
  id = "ecfg_xxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The ID of the Edge Config that the schema should be for.

### Optional

- `team_id` (String) The ID of the team the Edge Config should exist under. Required when configuring a team resource if a default team has not been set in the provider.

### Read-Only

- `definition` (String) A JSON schema that will be used to validate data in the Edge Config.
