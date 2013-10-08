package main
import (
  "code.google.com/p/goauth2/oauth"
  "github.com/google/go-github/github"
  "fmt"
  "os"
)

func main() {
  t := &oauth.Transport{
    Token: &oauth.Token{AccessToken: os.Getenv("GITHUB_KEY")},
  }

  client := github.NewClient(t.Client())

  // list all repositories for the authenticated user
  repos, _, err := client.Issues.List(true, nil)
  if err != nil {
    fmt.Printf("error: %v\n\n", err)
  } else {
    fmt.Printf("%v\n\n", github.Stringify(repos))
  }
}