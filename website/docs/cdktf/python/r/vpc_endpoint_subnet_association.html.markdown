---
subcategory: "VPC (Virtual Private Cloud)"
layout: "aws"
page_title: "AWS: aws_vpc_endpoint_subnet_association"
description: |-
  Provides a resource to create an association between a VPC endpoint and a subnet.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_vpc_endpoint_subnet_association

Provides a resource to create an association between a VPC endpoint and a subnet.

~> **NOTE on VPC Endpoints and VPC Endpoint Subnet Associations:** Terraform provides
both a standalone VPC Endpoint Subnet Association (an association between a VPC endpoint
and a single `subnet_id`) and a [VPC Endpoint](vpc_endpoint.html) resource with a `subnet_ids`
attribute. Do not use the same subnet ID in both a VPC Endpoint resource and a VPC Endpoint Subnet
Association resource. Doing so will cause a conflict of associations and will overwrite the association.

## Example Usage

Basic usage:

```python
# Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.vpc_endpoint_subnet_association import VpcEndpointSubnetAssociation
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        VpcEndpointSubnetAssociation(self, "sn_ec2",
            subnet_id=sn.id,
            vpc_endpoint_id=ec2.id
        )
```

## Argument Reference

The following arguments are supported:

* `vpc_endpoint_id` - (Required) The ID of the VPC endpoint with which the subnet will be associated.
* `subnet_id` - (Required) The ID of the subnet to be associated with the VPC endpoint.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The ID of the association.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `delete` - (Default `10m`)

## Import

VPC Endpoint Subnet Associations can be imported using `vpc_endpoint_id` together with `subnet_id`,
e.g.,

```
$ terraform import aws_vpc_endpoint_subnet_association.example vpce-aaaaaaaa/subnet-bbbbbbbbbbbbbbbbb
```

<!-- cache-key: cdktf-0.17.1 input-650992d61a270b5eff0edcd58e54247cd24d2eec19f5a8578dc025249b569472 -->