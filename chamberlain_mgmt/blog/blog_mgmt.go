package blog

type (
	Repository struct {
		RepoPath string
		RepoName string
	}
	Blogs struct {
		WorkPath  string
		BlogRepos []Repository
	}
	BlogsHandler interface {
		CleanWorkSpace() error
		CloneBlogsRepos() error
		ConvertBlogs2HTML() error
		GenerateTreeFile() error
	}
)

var blackDirNames = []string{"img"}

func (blogs *Blogs) CleanWorkSpace() error {
	return nil
}

func (blogs *Blogs) CloneBlogsRepos() error {
	return nil
}

func (blogs *Blogs) ConvertBlogs2HTML() error {
	return nil
}

func (blogs *Blogs) GenerateTreeFile() error {
	return nil
}