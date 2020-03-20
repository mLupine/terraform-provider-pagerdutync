package pagerdutync

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccPagerDutyTeam_import(t *testing.T) {
	team := fmt.Sprintf("tf-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPagerDutyTeamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPagerDutyTeamConfig(team),
			},

			{
				ResourceName:      "pagerduty_team.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}