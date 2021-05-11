package blog

import (
	"chamberlain_mgmt/log"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type (
	Article struct {
		FilePath        string
		FromFileName    string
		FromFileContent string
		ToFileName      string
		ToFileContent   string
		FilePermission  os.FileMode
	}
	ArticleConverter interface {
		Convert2HTML() error
		SaveToFile() error
		DeleteFromFile() error
	}
)

func (article *Article) Convert2HTML() error {
	filename := article.FilePath + string(os.PathSeparator) + article.FromFileName
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	content := template.HTML(blackfriday.Run(f))
	convertHtml := strings.ReplaceAll((string)(content), ".md", ".html")

	article.ToFileContent = convertHtml
	log.Debug(convertHtml)
	return nil
}

func (article *Article) SaveToFile() error {
	return ioutil.WriteFile(article.FilePath+string(os.PathSeparator)+article.ToFileName,
		[]byte(article.ToFileContent), article.FilePermission)
}

func (article *Article) DeleteFromFile() error {
	return os.Remove(article.FilePath + string(os.PathSeparator) + article.FromFileName)
}