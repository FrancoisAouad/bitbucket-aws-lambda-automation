package aws

import (
	"example/enums"
	"example/lib"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var logger = lib.Log

func GetAWSSecrets(secretName string, region string) *secretsmanager.GetSecretValueOutput {
	var bitbucketToken = enums.BITBUCKET_API_TOKEN
	fmt.Println(bitbucketToken)
	sess := session.Must(session.NewSession())
	svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))
	result, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secretName})
	if err != nil {
		logger.WithError(err).Error("Failed to get credentials from AWS Secret Manager.")
		os.Exit(1)
	}
	fmt.Println(*result.SecretString)
	return result
}
