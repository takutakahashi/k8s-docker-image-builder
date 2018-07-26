package github

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"os"
)

func Clone(repo string) string {
	// クローンしてくる
	os.RemoveAll(repo)
	_, err := git.PlainClone(repo, false, &git.CloneOptions{
		URL:      fmt.Sprintf("https://github.com/%s", repo),
		Progress: os.Stdout,
	})
	if err != nil {
		log.Print(err)
		panic(err)
	}
	return "repo"
}
