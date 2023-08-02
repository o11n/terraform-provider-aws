package identitystore

import (
	"context"
	"regexp"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_identitystore_groups")
func DataSourceGroups() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceGroupsRead,

		Schema: map[string]*schema.Schema{
			"identity_store_id": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.All(
					validation.StringLenBetween(1, 64),
					validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9-]*$`), "must match [a-zA-Z0-9-]"),
				),
			},
			"descriptions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"display_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"id": {
								Type:     schema.TypeString,
								Computed: true,
							},
							"issuer": {
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
			},
			"group_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

const (
	DSNameGroups = "Groups Data Source"
)

func dataSourceGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).IdentityStoreClient(ctx)

	identityStoreID := d.Get("identity_store_id").(string)

	groups := []types.Group{}

	var nextToken *string
	for {
		input := &identitystore.ListGroupsInput{
			IdentityStoreId: aws.String(identityStoreID),
			NextToken:       nextToken,
		}

		output, err := conn.ListGroups(ctx, input)
		if err != nil {
			return create.DiagError(names.IdentityStore, create.ErrActionReading, DSNameGroup, identityStoreID, err)
		}

		groups = append(groups, output.Groups...)

		if output.NextToken == nil {
			break
		}
		nextToken = output.NextToken
	}

	d.SetId(identityStoreID)

	sort.Slice(groups, func(i, j int) bool {
		return aws.ToString(groups[i].DisplayName) < aws.ToString(groups[j].DisplayName)
	})

	var descriptions, displayNames, groupIDs []string
	var externalIDs []any
	for _, group := range groups {
		descriptions = append(descriptions, aws.ToString(group.Description))
		displayNames = append(displayNames, aws.ToString(group.DisplayName))
		externalIDs = append(externalIDs, flattenExternalIds(group.ExternalIds))
		groupIDs = append(groupIDs, aws.ToString(group.GroupId))
	}

	d.Set("descriptions", descriptions)
	d.Set("display_names", displayNames)
	if err := d.Set("external_ids", externalIDs); err != nil {
		return create.DiagError(names.IdentityStore, create.ErrActionSetting, DSNameGroup, identityStoreID, err)
	}
	d.Set("group_ids", groupIDs)

	return nil
}
