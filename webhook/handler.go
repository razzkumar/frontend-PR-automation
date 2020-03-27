package webhook

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/razzkumar/sfebuild-tool/deployment"
	"github.com/razzkumar/sfebuild-tool/github"
	"github.com/razzkumar/sfebuild-tool/logger"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {

		PREvent, err := gh.Clone(c)

		commitHash := PREvent.PullRequest.GetHead().GetSHA()
		if err != nil {
			logger.FailOnError(err, "Cant't clone")
		}

		//Build fronted app
		build()

		//Deploy to s3
		fmt.Println("deploying")

		data := deployment.Data{
			BucketName: commitHash[0:7] + ".signoi",
		}

		err = deployment.Deploy(data)

		if err != nil {
			logger.FailOnError(err, "Unable to deploy")
		}

		url := deployment.GetURL(data.BucketName)

		gh.Comment(PREvent, url)

		c.JSON(http.StatusOK, gin.H{
			"Name":    "ghWebHook",
			"message": "looks ok",
		})

	}
}

func build() {

	cmd := exec.Command("yarn", "--cwd", "./repo")

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Info(string(output))
		logger.Info(err.Error())
	}
	fmt.Println(string(output))
	cmd = exec.Command("yarn", "--cwd", "./repo", "build")

	output, err = cmd.CombinedOutput()

	fmt.Println("------------build-----------------")

	if err != nil {
		logger.Info(string(output))
		logger.Info(err.Error())
	}

	fmt.Println(string(output))
}
