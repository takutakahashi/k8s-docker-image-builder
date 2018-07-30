package github

import (
	"fmt"
	"github.com/rs/xid"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	gitssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"io/ioutil"
	"log"
	"os"
)

func Clone(repo string) string {
	sshKey, err := ioutil.ReadFile("/root/.ssh/id_rsa")
	signer, err := ssh.ParsePrivateKey([]byte(sshKey))
	auth := &gitssh.PublicKeys{User: "git", Signer: signer}
	repoPath := repo + xid.New().String()
	_, err = git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      fmt.Sprintf("ssh://git@github.com/%s", repo),
		Progress: os.Stdout,
		Auth:     auth,
	})
	if err != nil {
		log.Print(err)
		panic(err)
	}
	return repoPath
}
