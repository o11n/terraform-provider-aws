package identitystore_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/identitystore"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccIdentityStoreGroupsDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)

	namePrefix := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	names := []string{
		fmt.Sprintf("%s-%d", namePrefix, 1),
		fmt.Sprintf("%s-%d", namePrefix, 2),
		fmt.Sprintf("%s-%d", namePrefix, 3),
	}

	dataSourceName := "data.aws_identitystore_groups.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			testAccPreCheckSSOAdminInstances(ctx, t)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGroupsDataSourceConfig_basicList(names[0], names[1], names[2]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "descriptions.#", "3"),
					resource.TestCheckResourceAttr(dataSourceName, "descriptions.0", "Acceptance Test 1"),
					resource.TestCheckResourceAttr(dataSourceName, "descriptions.1", "Acceptance Test 2"),
					resource.TestCheckResourceAttr(dataSourceName, "descriptions.2", "Acceptance Test 3"),

					resource.TestCheckResourceAttr(dataSourceName, "display_names.#", "3"),
					resource.TestCheckResourceAttr(dataSourceName, "display_names.0", names[0]),
					resource.TestCheckResourceAttr(dataSourceName, "display_names.1", names[1]),
					resource.TestCheckResourceAttr(dataSourceName, "display_names.2", names[2]),

					resource.TestCheckResourceAttr(dataSourceName, "external_ids.#", "0"),
					resource.TestCheckResourceAttr(dataSourceName, "group_ids.#", "3"),
				),
			},
		},
	})
}

func testAccGroupsDataSourceConfig_base(name1, name2, name3 string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_identitystore_group" "test_1" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = %[1]q
  description       = "Acceptance Test 1"
}

resource "aws_identitystore_group" "test_2" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = %[2]q
  description       = "Acceptance Test 2"
}
 
resource "aws_identitystore_group" "test_3" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = %[3]q
  description       = "Acceptance Test 3"
}
`, name1, name2, name3)
}

func testAccGroupsDataSourceConfig_basicList(name1, name2, name3 string) string {
	return acctest.ConfigCompose(
		testAccGroupsDataSourceConfig_base(name1, name2, name3),
		`
data "aws_identitystore_groups" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  depends_on = [aws_identitystore_group.test_1, aws_identitystore_group.test_2, aws_identitystore_group.test_3]
}
`,
	)
}
