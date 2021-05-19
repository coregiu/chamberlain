package blog

import (
	"chamberlain_mgmt/config"
	"fmt"
	"testing"
)

func TestBlogs_Should_generate_books_When_invoke_blog_handler(t *testing.T) {
	blogs := new(Blogs)
	blogs.WorkPath = "/giu/chamberlain/books"
	repos := make([]*config.BlogRepository, 2)
	blogs.BlogRepos = &repos
	repoA := new(config.BlogRepository)
	repoA.RepoName = "Philosophy"
	repoA.RepoPath = "git@gitee.com:regiu/philosophy.git"

	repoB := new(config.BlogRepository)
	repoB.RepoName = "Technology"
	repoB.RepoPath = "git@gitee.com:regiu/summary.git"

	repos[0] = repoA
	repos[1] = repoB

	err := blogs.CleanWorkSpace()
	if err != nil {
		fmt.Println(err.Error())
		t.Error(err.Error())
	}
	err = blogs.HandleBlogs()
	if err != nil {
		fmt.Println(err.Error())
		t.Error(err.Error())
	}
}