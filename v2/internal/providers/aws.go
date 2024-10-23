package providers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/datadog/stratus-red-team/v2/internal/utils"
	"github.com/google/uuid"
	"log"
	"os"
)

type AWSProvider struct {
	awsConfig           *aws.Config
	UniqueCorrelationId uuid.UUID // unique value injected in the user-agent, to differentiate Stratus Red Team executions
}

func NewAWSProvider(uuid uuid.UUID) *AWSProvider {
	cfg, err := config.LoadDefaultConfig(context.Background(), utils.CustomUserAgentApiOptions(uuid))
	if err != nil {
		log.Fatalf("unable to load AWS configuration, %v", err)
	}
	return &AWSProvider{UniqueCorrelationId: uuid, awsConfig: &cfg}
}
func (m *AWSProvider) GetConnection() aws.Config {
	return *m.awsConfig
}

func (m *AWSProvider) IsAuthenticatedAgainstAWS() bool {
	// Note: Explicitly setting AWS_REGION/AWS_DEFAULT_REGION is not strictly required for the AWS SDK to work, but it is necessary for Terraform
	// If it's not set, we get a user-unfriendly error such as the one describe at https://github.com/DataDog/stratus-red-team/issues/506
	if os.Getenv("AWS_REGION") == "" && os.Getenv("AWS_DEFAULT_REGION") == "" {
		return false
	}

	return true
}
