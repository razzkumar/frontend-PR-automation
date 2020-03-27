package deployment

import "github.com/razzkumar/sfebuild-tool/utils"

func GetURL(bucket string) string {
	region := utils.LoadEnv("AWS_REGION")
	url := "http://" + bucket + ".s3-website." + region + ".amazonaws.com/"
	return url
}
