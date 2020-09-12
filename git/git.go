package git

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/tcnksm/go-gitconfig"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"

	"github.com/integralist/go-findroot/find"
)

// Repository repo handle
type Repository struct {
	Repository *git.Repository
}

// GetRootDir returns the root directory of the git repository
func GetRootDir() (string, error) {
	dir, err := find.Repo()
	return dir.Path, err

}

// IsGitRepository returns true if the directory is a Git repository
func IsGitRepository(dir string) bool {
	_, err := git.PlainOpen(dir)

	return err == nil
}

// GetTags returns a list of git tags
func (r *Repository) GetTags() []string {
	tags := []string{}
	tagRefs, _ := r.Repository.Tags()

	err := tagRefs.ForEach(func(t *plumbing.Reference) error {
		tags = append(tags, strings.Split(t.String(), "/")[2])
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return tags
}

func getUsername() string {
	username, err := gitconfig.Username()
	if err != nil {
		log.Fatal(err)
	}

	return username
}

func getEmail() string {
	email, err := gitconfig.Email()
	if err != nil {
		log.Fatal(err)
	}

	return email
}

func defaultSignature() *object.Signature {
	return &object.Signature{
		Name:  getUsername(),
		Email: getEmail(),
		When:  time.Now(),
	}
}

// NewRepository returns the repository handle
func NewRepository(dir string) (*Repository, error) {
	r, e := git.PlainOpen(dir)
	return &Repository{r}, e
}

// AddTag adds a tag
func (r *Repository) AddTag(tag, message string) error {
	opts := &git.CreateTagOptions{
		Tagger:  defaultSignature(),
		Message: message,
	}

	head, err := r.Repository.Head()
	if err != nil {
		return err
	}

	_, err = r.Repository.CreateTag(tag, head.Hash(), opts)
	if err != nil {
		return err
	}

	return nil
}

// PushTags pushes to the specified remotre
func (r *Repository) PushTags(remote string) error {
	po := &git.PushOptions{
		RemoteName: remote,
		Progress:   os.Stdout,
		RefSpecs:   []config.RefSpec{config.RefSpec("refs/tags/*:refs/tags/*")},
	}

	err := r.Repository.Push(po)
	if err != nil {
		return err
	}

	return nil
}
