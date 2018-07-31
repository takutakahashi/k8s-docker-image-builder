package github

import (
	"fmt"
	"github.com/rs/xid"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"io/ioutil"
	"log"
	"os"
)

func Clone(repo string) string {
	sshKey, err := ioutil.ReadFile("/root/.ssh/id_rsa")
	if err != nil {
		log.Print(err)
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey([]byte(sshKey))
	if err != nil {
		log.Print(err)
		panic(err)
	}
	auth := &gitssh.PublicKeys{User: "git", Signer: signer}
	repoPath := repo + xid.New().String()
	_, err = git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:           fmt.Sprintf("ssh://git@github.com/%s", repo),
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", "master")),
		SingleBranch:  true,
		Progress:      os.Stdout,
		Auth:          auth,
	})
	if err != nil {
		log.Print(err)
		panic(err)
	}
	return repoPath
}
