package gh

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	//"github.com/go-git/go-git/v5/plumbing"
	"github.com/google/go-github/v30/github"
	"github.com/razzkumar/sfebuild-tool/logger"
	"github.com/razzkumar/sfebuild-tool/utils"
	ghhttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func Clone(c *gin.Context) (*github.PullRequestEvent, error) {

	accssToken := utils.LoadEnv("GH_ACCSS_TOKEN")
	prEvent, err := getRepo(c)
	fmt.Println(accssToken)
	if err != nil {
		logger.FailOnError(err, "Error While parsing gh-event")
	}

	fmt.Println("+++++++++++++Clone++++++++++++++++++++")
	repo, err := git.PlainClone("./repo", false, &git.CloneOptions{
		//URL: prEvent.GetRepo().GetCloneURL(),
		URL:      "https://github.com/razzkumar/ftodo",
		Progress: os.Stderr,
		Auth: &ghhttp.BasicAuth{
			Username: "razzkumar",
			Password: accssToken,
		},
	})

	if err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir("./repo")

	if err != nil {
		return nil, err
	}
	fmt.Println("-------------files----------------------")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	w, err := repo.Worktree()
	if err != nil {
		return nil, err
	}
	fmt.Println(w)
	//// ... checking out to commit
	//logger.Info("git checkout " + commitHash)
	//err = w.Checkout(&git.CheckoutOptions{
	//Hash: plumbing.NewHash(commitHash),
	//})

	//if err != nil {
	//return nil, err
	//}

	//// ... retrieving the commit being pointed by HEAD, it shows that the
	//// repository is pointing to the giving commit in detached mode
	//logger.Info("git show-ref --head HEAD")
	//ref, err := repo.Head()
	//if err != nil {
	//return nil, err
	//}
	//fmt.Println(ref.Hash())

	return prEvent, nil
}

// Parse Hook Body response and return PullRequestEvent
func getRepo(c *gin.Context) (*github.PullRequestEvent, error) {

	signature := c.GetHeader("X-Hub-Signature")

	if signature == "" {
		return nil, fmt.Errorf("Invalid webhook Signature")
	}

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		return nil, fmt.Errorf("Unable to parse requestBody: %s", err)
	}

	defer c.Request.Body.Close()

	// Check the event type. Make sure just a PR
	// We need to get the event from the Request object as Gin in the middle
	// does some normalization that breaks this particular header name.

	event := c.Request.Header.Get("X-GitHub-Event")

	parsedData, err := github.ParseWebHook(event, body)

	if err != nil {
		return nil, fmt.Errorf("Unable to Read requestBody: %s", err)
	}

	switch event := parsedData.(type) {
	case *github.PullRequestEvent:
		return event, nil
	default:
		return nil, fmt.Errorf("The event type is not supported")
	}
}
