package main

import (
	//"fmt"
	//"io"
	"context"
	"log"
	"net/http"
	"os"
	"github.com/google/go-github/github"
	"strings"
	"golang.org/x/oauth2"
	"flag"
)

var (
	owner string
	repoName string
	number int
	token string
	secret string
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if secret == "" {
		log.Fatal("-secret or GITHUB_SECRET required")
	}
	secretBytes := []byte(secret)
	payload, err := github.ValidatePayload(r, secretBytes)
	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		log.Printf("could not parse webhook: err=%s\n", err)
		return
	}

	switch e := event.(type) {
	case *github.PullRequestReviewEvent:
		log.Println("Pull Request Review")
		// this is a commit push, do something with it
	case *github.PullRequestReviewCommentEvent:
		log.Println("PR Comment: ")
		log.Println(e.GetComment())
		// this is a pull request, do something with it
	case *github.IssueCommentEvent:
		// Inclueds comments on a PR
		repo := e.GetRepo()
		repoName = repo.GetName()
		owner = repo.GetOwner().GetLogin()
		number = e.GetIssue().GetNumber()
		//log.Println(e.GetNumber())
		parseComment(*e.GetComment())
	default:
		log.Printf("unknown event type %s\n", github.WebHookType(r))
		return
	}
}

// Checks to see if /retest is in the comment
func parseComment(com github.IssueComment) {
	body := com.GetBody()
	//log.Println(body)
	if strings.Contains(body, "/retest") {
		log.Println("Triggering Retest")
		writeComment()
	}
}

// writes Retest Initiated after 
func writeComment() {
	//getting the github client
	if token == "" {
		log.Fatal("-token or GITHUB_TOKEN required")
	}


	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)


	commentStr := "Retest Initiated" // temporary
	issueComment := &github.IssueComment{Body: &commentStr}
	// ToDo: get the number from the issue comment recieved by the server
	var err error // needs to be declared for the next line
	issueComment, _, err = client.Issues.CreateComment(context.Background(), owner, repoName, number, issueComment)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	token = os.Getenv("GITHUB_TOKEN")
	secret = os.Getenv("GITHUB_SECRET")
    log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}