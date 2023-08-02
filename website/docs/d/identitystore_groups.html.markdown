---
subcategory: "SSO Identity Store"
layout: "aws"
page_title: "AWS: aws_identitystore_groups"
description: |-
  List all Identity Store Groups
---

# Data Source: aws_identitystore_groups

Use this data source to list all groups in an Identity Store.

## Example Usage

### Basic Usage

```terraform
data "aws_identitystore_groups" "this" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.example.identity_store_ids)[0]
}
```

## Argument Reference

The following arguments are required:

* `identity_store_id` - (Required) Identity Store ID associated with the Single Sign-On Instance.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `descriptions` - A list of strings containing the description of a group.
* `display_names` - A list of strings containing the display name of a group.
* `external_ids` - A list of list of identifiers issued to the group by an external identity provider.
    * `id` - The identifier issued to the group by an external identity provider.
    * `issuer` - The issuer for an external identifier.
* `group_ids` - A list of strings containing the group IDs.

Note that the indexes of `descriptions`, `display_names`, `external_ids`, and `group_ids` correspond.
