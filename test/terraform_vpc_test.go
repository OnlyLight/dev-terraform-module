package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformVPC(t *testing.T) {
	t.Parallel()

	// Pick AWS region
	// awsRegion := "ap-southeast-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../vpc",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.WorkspaceSelectOrNew(t, terraformOptions, "dev")

	terraform.InitAndPlan(t, terraformOptions)

	// publicSubnetZone1Id := terraform.Output(t, terraformOptions, "public_zone1_id")
	// publicSubnetZone2Id := terraform.Output(t, terraformOptions, "public_zone2_id")
	// vpcID := terraform.Output(t, terraformOptions, "main_vpc_id")

	// subnets := aws.GetSubnetsForVpc(t, vpcID, awsRegion)

	// require.Equal(t, 4, len(subnets))

	// assert.True(t, aws.IsPublicSubnet(t, publicSubnetZone1Id, awsRegion))
	// assert.True(t, aws.IsPublicSubnet(t, publicSubnetZone2Id, awsRegion))
}
