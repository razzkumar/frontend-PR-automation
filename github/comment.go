package gh

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v30/github"
	"github.com/razzkumar/sfebuild-tool/utils"
	"golang.org/x/oauth2"
)

func Comment(prEvent *github.PullRequestEvent, url string) {
	repo := prEvent.GetRepo().Name
	fmt.Println(repo)

	owner := prEvent.GetRepo().GetOwner().GetLogin()
	fmt.Println(owner)

	ctx := context.Background()

	prEvent.PullRequest.GetComments()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: utils.LoadEnv("GH_ACCSS_TOKEN")},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	num := prEvent.GetNumber()

	comment := "Visit: " + url

	pullRequestReviewRequest := &github.PullRequestReviewRequest{Body: &comment, Event: github.String("COMMENT")}

	//client.PullRequests.CreateComment(ctx, owner, repo, num, pullRequestReviewRequest)
	pullRequestReview, _, err := client.PullRequests.CreateReview(ctx, owner, *repo, num, pullRequestReviewRequest)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("github-commenter: Created GitHub PR Review comment", pullRequestReview.ID)
}
