package blog

import (
	"chamberlain_mgmt/log"
	"github.com/gin-gonic/gin/json"
	"io/ioutil"
)

type (
	RootTree struct {
		Root *[]*TreeNode `json:"root"`
	}
	TreeNode struct {
		Key      string       `json:"key"`
		Label    string       `json:"label"`
		Icon     string       `json:"icon"`
		Data     string       `json:"data"`
		Link     string       `json:"link"`
		Children *[]*TreeNode `json:"children"`
	}
	TreeHandle interface {
		CreateRootTree(size int)
		AppendRootChildren(treeNode *TreeNode, index int)
		CreateChildTree(size int)
		AppendChildren(treeNode *TreeNode, index int)
		CleanNilTree()
		TrimNodes(index int)
		GenerateTreeFile(filePath string)
	}
)

func (rootTree *RootTree) CreateRootTree(size int) {
	treeNodes := make([]*TreeNode, size)
	rootTree.Root = &treeNodes
}

func (rootTree *RootTree) AppendRootChildren(treeNode *TreeNode, index int) {
	(*rootTree.Root)[index] = treeNode
}

func (currentTree *TreeNode) CreateChildTree(size int) {
	node := make([]*TreeNode, size)
	currentTree.Children = &node
}

func (currentTree *TreeNode) AppendChildren(treeNode *TreeNode, index int) {
	(*currentTree.Children)[index] = treeNode
}

func (currentTree *TreeNode) CleanNilTree() {
	if len(*currentTree.Children) == 0 {
		return
	}
	if (*currentTree.Children)[0] == nil {
		currentTree.Children = nil
	}
}

func (currentTree *TreeNode) TrimNodes(index int) {
	newTree := (*currentTree.Children)[0:index]
	currentTree.Children = &newTree
}
func (rootTree *RootTree) GenerateTreeFile(filePath string) {
	jsonByte, err := json.Marshal(rootTree)
	if err != nil {
		log.Error("failed to convert root tree to json for %s", err.Error())
		return
	}
	log.Info("books = %s", string(jsonByte))
	err = ioutil.WriteFile(filePath, jsonByte, 0777)
	if err != nil {
		log.Error("failed to save root tree json file for %s", err.Error())
	}
}
