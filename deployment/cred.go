package deployment

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/razzkumar/sfebuild-tool/logger"
	"github.com/razzkumar/sfebuild-tool/utils"
)

// GetSession returns Session for AWS.
func GetSession() *session.Session {

	region := utils.LoadEnv("AWS_REGION")
	accssKey := utils.LoadEnv("AWS_ACCESS_KEY_ID")
	secrectKey := utils.LoadEnv("AWS_SECRET_ACCESS_KEY")
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(accssKey, secrectKey, ""), // token empty
		},
	})

	if err != nil {
		logger.FailOnError(err, "Unable to connect connect to AWS.")
	}

	return sess
}
