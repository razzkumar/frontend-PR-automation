package deployment

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/razzkumar/sfebuild-tool/logger"
	"github.com/razzkumar/sfebuild-tool/utils"
)

// Data contains the data needed to deploy to S3 bucket
type Data struct {
	BucketName string
	DistDir    string
}

// Deploy to S3 bucket
func Deploy(data Data) error {
	data.DistDir = "./repo/build"

	sess := GetSession()
	svc := s3.New(sess)
	err := CreateBucket(data.BucketName, svc)

	if err != nil {
		logger.Info(err.Error())
	}

	uploader := s3manager.NewUploader(sess)

	fileList := []string{}

	filepath.Walk(data.DistDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			fileList = append(fileList, path)

			return nil
		})

	// Loop through every file and uplaod to s3
	for _, file := range fileList {
		f, _ := os.Open(file)

		key := strings.TrimPrefix(file, data.DistDir)
		key = strings.Replace(key, "repo/build", "", -1)
		fileContentType := utils.GetFileType(file)

		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket:      aws.String(data.BucketName),
			Key:         aws.String(key),
			ContentType: aws.String(fileContentType),
			Body:        f,
		})

		if err != nil {
			return err
		}
		fmt.Println("Uploading... " + key)
	}

	fmt.Println("\n\n" + strconv.Itoa(len(fileList)) + " Files Uploaded Successfully. 🎉 🎉 🎉")
	fmt.Println("removeing filse")
	os.RemoveAll("./repo")
	return nil
}
