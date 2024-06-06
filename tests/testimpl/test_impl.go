package common

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/require"
)

func TestAcmCertExists(t *testing.T, ctx types.TestContext) {
	acmClient := acm.NewFromConfig(GetAWSConfig(t))
	certArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "certificate_arn")
	certStatus := terraform.Output(t, ctx.TerratestTerraformOptions(), "certificate_status")
	certRenewal := terraform.Output(t, ctx.TerratestTerraformOptions(), "renewal_eligibility")

	t.Run("TestDoesCertExist", func(t *testing.T) {
		output, err := acmClient.DescribeCertificate(context.TODO(), &acm.DescribeCertificateInput{CertificateArn: &certArn})
		if err != nil {
			t.Errorf("Error getting certificate: %v", err)
		}

		require.Equal(t, certArn, *output.Certificate.CertificateArn, "Expected certificate ARN to match")
		require.Equal(t, certStatus, string(output.Certificate.Status), "Expected certificate status to match")
		require.Equal(t, certRenewal, string(output.Certificate.RenewalEligibility), "Expected certificate renewal eligibility to match")
	})
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
