package awsclient

import (
	"photon-server/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// CreateAWSSession create AWS session
func CreateAWSSession() (*session.Session, error) {

	cfg := config.NewConfig()
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())
	creds := credentials.NewStaticCredentials(cfg.AwsCredentials.AwsAccessKey, cfg.AwsCredentials.AwsSecretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String("ap-northeast-1")},
	)
	return sess, err
}
