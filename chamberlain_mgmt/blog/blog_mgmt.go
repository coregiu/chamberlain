package blog

import (
	"chamberlain_mgmt/config"
	"chamberlain_mgmt/log"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

type (
	Blogs struct {
		WorkPath    string
		BlogRepos   *[]*config.BlogRepository
		innerParams *innerParams
	}
	BlogsHandler interface {
		CleanWorkSpace() error
		HandleBlogs() error
		CloneBlogsRepos()
		ConvertBlogs2HTML(rootTree *RootTree)
	}
	innerParams struct {
		cloneQueue   *chan *config.BlogRepository
		convertQueue *chan *config.BlogRepository
	}
)

// black list for files not need display.
var blackFileNames = []string{"img", ".git", "archive"}
// global synchronized tool for avoid repeat handles.
var waitGroup = &sync.WaitGroup{}

func (blogs *Blogs) CleanWorkSpace() error {
	commandStr := "cd " + blogs.WorkPath + " && rm -rf *"
	cmd := exec.Command("/bin/bash", "-c", commandStr)
	return cmd.Run()
}

func (blogs *Blogs) HandleBlogs() error {
	blogs.innerParams = &innerParams{}
	cloneQueue := make(chan *config.BlogRepository)
	blogs.innerParams.cloneQueue = &cloneQueue
	convertQueue := make(chan *config.BlogRepository)
	blogs.innerParams.convertQueue = &convertQueue
	waitGroup.Add(len(*blogs.BlogRepos))

	rootTree := new(RootTree)
	rootTree.CreateRootTree(len(*blogs.BlogRepos))

	go blogs.CloneBlogsRepos()
	go blogs.ConvertBlogs2HTML(rootTree)

	for iLoop, repo := range *blogs.BlogRepos {
		log.Info("loop = %d", iLoop)
		*blogs.innerParams.cloneQueue <- repo
	}

	waitGroup.Wait()
	close(*blogs.innerParams.cloneQueue)
	close(*blogs.innerParams.convertQueue)

	rootTree.GenerateTreeFile(blogs.WorkPath + string(os.PathSeparator) + "books.json")
	return nil
}

func (blogs *Blogs) CloneBlogsRepos() {
	for {
		repo, isOpen := <-*blogs.innerParams.cloneQueue
		if !isOpen {
			log.Warn("cloneQueue was closed")
			break
		}
		log.Info("begin to clone the repo %s, %s", repo.RepoPath, repo.RepoName)
		cmdStr := "cd " + blogs.WorkPath + " && git clone " + repo.RepoPath + " " + repo.RepoName

		cmd := exec.Command("/bin/bash", "-c", cmdStr)
		err := cmd.Run()
		if err != nil {
			log.Error("Failed to clone the repo %s", err.Error())
		}
		*blogs.innerParams.convertQueue <- repo
		time.Sleep(time.Second)
	}
}

func (blogs *Blogs) ConvertBlogs2HTML(rootTree *RootTree) {
	for iLoop := 0; ; iLoop++ {
		repo, isOpen := <-*blogs.innerParams.convertQueue
		if !isOpen {
			log.Warn("convertQueue was closed")
			break
		}
		repoNode := generateDirNode(blogs.WorkPath, repo.RepoName)
		rootTree.AppendRootChildren(repoNode, iLoop)
		convertDirs(blogs.WorkPath, repo.RepoName, repoNode)
		waitGroup.Done()
		time.Sleep(time.Second)
	}
}

func convertDirs(dirPath string, fileName string, parentTreeNode *TreeNode) {
	repoPath := dirPath + string(os.PathSeparator) + fileName
	log.Info("begin to convert the repo %s", repoPath)
	files, _ := ioutil.ReadDir(repoPath)
	if len(files) == 0 {
		log.Warn("there is no file in the repo %s", repoPath)
		return
	}
	parentTreeNode.CreateChildTree(len(files))

	index := 0
	for iLoop, file := range files {
		log.Info("the %d's file %s", iLoop, file.Name())
		isInBlack := false
		for _, blackFile := range blackFileNames {
			if strings.Compare(file.Name(), blackFile) == 0 {
				isInBlack = true
				break
			}
		}
		if isInBlack {
			log.Info("%s was in black list.", file.Name())
			continue
		}
		if file.IsDir() {
			currentNode := generateDirNode(repoPath, file.Name())
			parentTreeNode.AppendChildren(currentNode, index)
			convertDirs(repoPath, file.Name(), currentNode)
		} else {
			fileNameWithoutSuffix, isMarkDownFile := getFileNameWithoutSuffix(file.Name())
			if !isMarkDownFile {
				continue
			}
			toFileName := fileNameWithoutSuffix + ".html"
			createAndConvertToFile(repoPath, file.Name(), toFileName)
			if strings.Compare(file.Name(), "README.md") == 0 {
				continue
			}
			currentNode := generateFileNode(repoPath, fileNameWithoutSuffix)
			parentTreeNode.AppendChildren(currentNode, index)
		}
		index++
	}
	// if there is no children, then clear the arrays;
	// or if no enough children as defined before(files length), then trim children, avoid null in json.
	if index == 0 {
		parentTreeNode.CleanNilTree()
	} else if index < len(files) {
		parentTreeNode.TrimNodes(index)
	}
}

func createAndConvertToFile(repoPath string, fromFileName string, toFileName string) {
	article := new(Article)
	article.FilePath = repoPath
	article.FilePermission = 0777
	article.FromFileName = fromFileName
	article.ToFileName = toFileName
	err := article.Convert2HTML()
	if err != nil {
		log.Error("failed to convert md to html %s, %s", article.FilePath, article.FromFileName)
	}
	err = article.SaveToFile()
	if err != nil {
		log.Error("failed to save to html file %s, %s", article.FilePath, article.FromFileName)
	}
	err = article.DeleteFromFile()
	if err != nil {
		log.Error("failed to delete md file %s, %s", article.FilePath, article.FromFileName)
	}
}

func getFileNameWithoutSuffix(fileName string) (string, bool) {
	if strings.LastIndex(fileName, ".md") != -1 {
		return string([]rune(fileName)[:utf8.RuneCountInString(fileName)-3]), true
	} else {
		return fileName, false
	}
}

func generateDirNode(parentPath string, dirName string) *TreeNode {
	node := new(TreeNode)
	node.Key = uuid.NewString()
	node.Label = dirName
	node.Data = dirName
	node.Icon = "pi pi-fw pi-inbox"
	node.Link = getSitePrefixPath(parentPath) + string(os.PathSeparator) + dirName + string(os.PathSeparator) + "README.html"
	return node
}

func generateFileNode(parentPath string, fileName string) *TreeNode {
	node := new(TreeNode)
	node.Key = uuid.NewString()
	node.Label = fileName
	node.Data = fileName
	node.Icon = "pi pi-fw pi-file"
	node.Link = getSitePrefixPath(parentPath) + string(os.PathSeparator) + fileName + ".html"
	return node
}

func getSitePrefixPath(path string) string {
	index := strings.Index(path, "/books")
	if index == -1 {
		return path
	}
	return "../.." + string([]rune(path)[index:utf8.RuneCountInString(path)])
}
